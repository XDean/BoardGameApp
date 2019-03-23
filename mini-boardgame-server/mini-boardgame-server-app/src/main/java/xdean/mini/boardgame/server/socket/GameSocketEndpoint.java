package xdean.mini.boardgame.server.socket;

import static xdean.jex.util.lang.ExceptionUtil.uncheck;

import java.io.IOException;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.CopyOnWriteArrayList;

import javax.inject.Inject;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.core.userdetails.User;
import org.springframework.stereotype.Component;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketSession;
import org.springframework.web.socket.handler.TextWebSocketHandler;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import io.reactivex.disposables.CompositeDisposable;
import io.reactivex.disposables.Disposable;
import io.reactivex.schedulers.Schedulers;
import io.reactivex.subjects.PublishSubject;
import io.reactivex.subjects.Subject;
import xdean.jex.log.Logable;
import xdean.mini.boardgame.server.model.CommonConstants;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.mybatis.mapper.GameMapper;
import xdean.mini.boardgame.server.security.TokenAuthProvider;
import xdean.mini.boardgame.server.service.UserDataService;

@Component
public class GameSocketEndpoint extends TextWebSocketHandler implements Logable, CommonConstants {

  private final Map<Integer, GameRoomSocketHandler> rooms = new ConcurrentHashMap<>();

  private @Autowired(required = false) List<GameSocketProvider> providers = Collections.emptyList();
  private @Inject ObjectMapper objectMapper;
  private @Inject GameMapper gameMapper;
  private @Inject UserDataService userService;
  private @Inject TokenAuthProvider tokenAuth;

  @Override
  public void afterConnectionEstablished(WebSocketSession session) throws Exception {
    GameRoomEntity room;
    try {
      int id = getRoomIdFromSession(session);
      room = gameMapper.findRoom(id);
    } catch (Exception ex) {
      trace("Fail to find room: " + session, ex);
      session.close(CloseStatus.BAD_DATA.withReason("Can't find room id"));
      return;
    }
    if (room != null) {
      trace(session.getRemoteAddress() + " connect to room " + room.getId());
      GameRoomSocketHandler roomHandler = rooms.computeIfAbsent(room.getId(), i -> new GameRoomSocketHandler(room));
      roomHandler.addSession(session);
    }
  }

  @Override
  public void afterConnectionClosed(WebSocketSession session, CloseStatus status) throws Exception {
    int id = getRoomIdFromSession(session);
    GameRoomSocketHandler room = rooms.get(id);
    if (room != null) {
      room.removeSession(session);
    }
  }

  @Override
  protected void handleTextMessage(WebSocketSession session, TextMessage message) throws Exception {
    int id = getRoomIdFromSession(session);
    GameRoomSocketHandler room = rooms.get(id);
    if (room != null) {
      room.handleMessage(session, message.getPayload());
    }
  }

  private int getRoomIdFromSession(WebSocketSession session) {
    String path = session.getUri().getPath();
    String idStr = path.substring(path.lastIndexOf('/') + 1);
    int id = Integer.parseInt(idStr);
    return id;
  }

  private class GameRoomSocketHandler {
    final GameRoomEntity room;
    final List<WebSocketSession> sessions = new CopyOnWriteArrayList<>();
    final Map<WebSocketSession, Subject<WebSocketEvent<JsonNode>>> subjects = new HashMap<>();
    final Map<WebSocketSession, Disposable> disposables = new HashMap<>();

    GameRoomSocketHandler(GameRoomEntity room) {
      this.room = room;
    }

    void addSession(WebSocketSession session) {
      sessions.add(session);
    }

    void initSession(WebSocketSession session) {
      Subject<WebSocketEvent<JsonNode>> messageSubject = PublishSubject.create();
      CompositeDisposable disposable = new CompositeDisposable();
      providers.forEach(p -> disposable.add(p.handle(session, room, messageSubject.observeOn(Schedulers.io()))
          .subscribeOn(Schedulers.io())
          .subscribe(e -> sendMessage(session, e),
              e -> error("Unhandled error happens: " + session, e))));
      subjects.put(session, messageSubject);
      disposables.put(session, disposable);
      trace("Websocket subscribe to provider: " + session);
    }

    void removeSession(WebSocketSession session) {
      sessions.remove(session);
      Subject<WebSocketEvent<JsonNode>> subject = subjects.remove(session);
      if (subject != null) {
        subject.onComplete();
      }
      disposables.remove(session);// TODO Dean, dispose or not
    }

    void handleMessage(WebSocketSession session, String message) {
      try {
        WebSocketEvent<JsonNode> event = objectMapper.readValue(message, new TypeReference<WebSocketEvent<JsonNode>>() {
        });
        if (session.getAttributes().get(AttrKey.ACCESS_TOKEN) == null) {
          if (event.topic.equals(SocketTopic.AUTHENTICATION)) {
            String token = event.attributes.getOrDefault(AttrKey.ACCESS_TOKEN, "").toString();
            Authentication authenticate;
            try {
              authenticate = tokenAuth.authenticate(token);
            } catch (AuthenticationException e) {
              sendMessage(session, WebSocketEvent.builder()
                  .id(event.getId())
                  .type(WebSocketSendType.SELF)
                  .topic(SocketTopic.BAD_CREDENTIAL)
                  .payload("Bad Credential: " + e.getLocalizedMessage())
                  .build());
              return;
            }
            User user = (User) authenticate.getPrincipal();
            Optional<UserEntity> ue = userService.findUserByUsername(user.getUsername());
            if (ue.isPresent()) {
              session.getAttributes().put(AttrKey.USER_ID, ue.get().getId());
              session.getAttributes().put(AttrKey.ACCESS_TOKEN, token);
              initSession(session);
              sendMessage(session, WebSocketEvent.builder()
                  .id(event.getId())
                  .type(WebSocketSendType.SELF)
                  .topic(SocketTopic.AUTHENTICATION)
                  .build());
            } else {
              error("An authed user not in db: " + user.getUsername());
              sendMessage(session, WebSocketEvent.builder()
                  .id(event.getId())
                  .type(WebSocketSendType.SELF)
                  .topic(SocketTopic.ERROR_TOPIC)
                  .payload("No such user, unexpected server error")
                  .build());
            }
          } else {
            sendMessage(session, WebSocketEvent.builder()
                .id(event.getId())
                .type(WebSocketSendType.SELF)
                .topic(SocketTopic.ERROR_TOPIC)
                .payload("The web socket should AUTHENTICATION first")
                .build());
          }
        } else if (event.topic.equals(SocketTopic.AUTHENTICATION)) {
          // TODO
        } else {
          Subject<WebSocketEvent<JsonNode>> subject = subjects.get(session);
          if (subject != null) {
            subject.onNext(event);
          }
        }
      } catch (IOException e) {
        sendMessage(session, WebSocketEvent.builder()
            .type(WebSocketSendType.SELF)
            .topic(SocketTopic.ERROR_TOPIC)
            .payload("Wrong format. All messages must be like this message.")
            .build());
        trace("Session send wrong message: " + session, e);
      }
    }

    void sendMessage(WebSocketSession session, WebSocketEvent<?> message) {
      try {
        String str = objectMapper.writeValueAsString(message);
        TextMessage msg = new TextMessage(str);
        if (message.type == WebSocketSendType.SELF) {
          if (session.isOpen()) {
            session.sendMessage(msg);
          }
        } else {
          sessions.stream().filter(s -> s.isOpen()).forEach(s -> uncheck(() -> s.sendMessage(msg)));
        }
      } catch (IOException e) {
        Subject<WebSocketEvent<JsonNode>> subject = subjects.get(session);
        if (subject != null) {
          subject.onError(e);
        } else {
          debug("Error happing after session close: " + session, e);
        }
      }
    }
  }
}
