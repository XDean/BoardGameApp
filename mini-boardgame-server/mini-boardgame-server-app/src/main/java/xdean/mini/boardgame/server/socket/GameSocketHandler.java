package xdean.mini.boardgame.server.socket;

import static xdean.jex.util.lang.ExceptionUtil.uncheck;

import java.io.IOException;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.CopyOnWriteArrayList;

import org.springframework.beans.factory.annotation.Autowired;
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
import xdean.mini.boardgame.server.model.GameConstants;
import xdean.mini.boardgame.server.model.GameRoom;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

@Component
public class GameSocketHandler extends TextWebSocketHandler implements Logable, GameConstants {

  private final Map<Integer, GameRoomSocketHandler> rooms = new ConcurrentHashMap<>();

  @Autowired(required = false)
  List<GameSocketProvider> providers = Collections.emptyList();

  ObjectMapper objectMapper = new ObjectMapper();

  @Override
  public void afterConnectionEstablished(WebSocketSession session) throws Exception {
    GameRoomEntity e = (GameRoomEntity) session.getAttributes().get(ROOM);
    if (e != null) {
      trace(session.getRemoteAddress() + " connect to room " + e.getId());
      GameRoomSocketHandler room = rooms.computeIfAbsent(e.getId(), i -> new GameRoomSocketHandler(e.getRoom()));
      room.addSession(session);
    }
  }

  @Override
  public void afterConnectionClosed(WebSocketSession session, CloseStatus status) throws Exception {
    GameRoomEntity e = (GameRoomEntity) session.getAttributes().get(ROOM);
    if (e != null) {
      GameRoomSocketHandler room = rooms.get(e.getId());
      if (room != null) {
        room.removeSession(session);
      }
    }
  }

  @Override
  protected void handleTextMessage(WebSocketSession session, TextMessage message) throws Exception {
    GameRoomEntity e = (GameRoomEntity) session.getAttributes().get(ROOM);
    if (e != null) {
      GameRoomSocketHandler room = rooms.get(e.getId());
      if (room != null) {
        room.handleMessage(session, message.getPayload());
      }
    }
  }

  private class GameRoomSocketHandler {
    final GameRoom room;
    final List<WebSocketSession> sessions = new CopyOnWriteArrayList<>();
    final Map<WebSocketSession, Subject<WebSocketEvent<JsonNode>>> subjects = new HashMap<>();
    final Map<WebSocketSession, Disposable> disposables = new HashMap<>();

    GameRoomSocketHandler(GameRoom room) {
      this.room = room;
    }

    void addSession(WebSocketSession session) {
      sessions.add(session);
      Subject<WebSocketEvent<JsonNode>> messageSubject = PublishSubject.create();
      CompositeDisposable disposable = new CompositeDisposable();
      providers.forEach(p -> disposable.add(p.handle(session, room, messageSubject.observeOn(Schedulers.io()))
          .subscribeOn(Schedulers.io())
          .subscribe(e -> sendMessage(session, e),
              e -> warn("Unhandled error happens: " + session, e))));
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
        Subject<WebSocketEvent<JsonNode>> subject = subjects.get(session);
        if (subject != null) {
          subject.onNext(event);
        }
      } catch (IOException e) {
        sendMessage(session, WebSocketEvent.builder()
            .type(WebSocketSendType.SELF)
            .topic(WebSocketEvent.ERROR_TOPIC)
            .payload(e.getMessage())
            .build());
        trace("Session send wrong message: " + session, e);
      }
    }

    void sendMessage(WebSocketSession session, WebSocketEvent<?> message) {
      try {
        String str = objectMapper.writeValueAsString(message);
        TextMessage msg = new TextMessage(str);
        if (message.type == WebSocketSendType.SELF) {
          session.sendMessage(msg);
        } else {
          sessions.forEach(s -> uncheck(() -> s.sendMessage(msg)));
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
