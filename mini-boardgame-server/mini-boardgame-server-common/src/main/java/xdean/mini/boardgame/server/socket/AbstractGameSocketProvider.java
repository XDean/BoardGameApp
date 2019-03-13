package xdean.mini.boardgame.server.socket;

import java.util.HashMap;
import java.util.Map;

import org.springframework.util.Assert;
import org.springframework.web.socket.WebSocketSession;

import com.fasterxml.jackson.databind.JsonNode;

import io.reactivex.Observable;
import io.reactivex.Observer;
import io.reactivex.disposables.Disposable;
import io.reactivex.subjects.PublishSubject;
import io.reactivex.subjects.Subject;
import lombok.AccessLevel;
import lombok.Builder;
import lombok.Value;
import lombok.experimental.FieldDefaults;
import xdean.jex.log.Logable;
import xdean.mini.boardgame.server.model.GlobalConstants;
import xdean.mini.boardgame.server.model.GlobalConstants.SocketTopic;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

public abstract class AbstractGameSocketProvider implements GameSocketProvider, Logable {

  @Value
  @FieldDefaults(level = AccessLevel.PUBLIC, makeFinal = true)
  @Builder(toBuilder = true)
  public static class SocketContext {
    WebSocketSession session;
    GameRoomEntity room;
    int userId;
    Observable<WebSocketEvent<JsonNode>> inputFlow;
    Observable<WebSocketEvent<?>> outputFlow;
    Observer<WebSocketEvent<?>> outputObserver;
  }

  Map<Integer, Subject<WebSocketEvent<?>>> playerSubjects = new HashMap<>();

  public void sendEvent(int playerId, WebSocketEvent<?> event) {
    Subject<WebSocketEvent<?>> subject = playerSubjects.get(playerId);
    if (subject != null) {
      subject.onNext(event);
    }
  }

  @Override
  public Observable<WebSocketEvent<?>> handle(WebSocketSession session, GameRoomEntity room,
      Observable<WebSocketEvent<JsonNode>> input) {
    Integer id = (Integer) session.getAttributes().get(GlobalConstants.AttrKey.USER_ID);
    Assert.notNull(id, "Authed user must have id");
    SocketContext context = SocketContext.builder()
        .session(session)
        .room(room)
        .userId(id)
        .inputFlow(input)
        .build();
    Subject<WebSocketEvent<?>> subject = playerSubjects.computeIfAbsent(id, i -> createOutputFlow(context));
    SocketContext processedContext = initFlow(context.toBuilder().outputFlow(subject).outputObserver(subject).build());
    subscribeInputFlow(processedContext);
    return processedContext.outputFlow;
  }

  protected Subject<WebSocketEvent<?>> createOutputFlow(SocketContext context) {
    return PublishSubject.create();
  }

  protected SocketContext initFlow(SocketContext context) {
    return context.toBuilder()
        .inputFlow(processInputFlow(context))
        .outputFlow(processOutputFlow(context))
        .build();
  }

  protected Observable<WebSocketEvent<JsonNode>> processInputFlow(SocketContext context) {
    return context.inputFlow;
  }

  protected Observable<WebSocketEvent<?>> processOutputFlow(SocketContext context) {
    return context.outputFlow;
  }

  protected Disposable subscribeInputFlow(SocketContext processedContext) {
    return processedContext.inputFlow.subscribe(e -> {
    },
        e -> {
          if (e instanceof RuntimeException) {
            debug("Unexpected error happens: " + processedContext.session, e);
            processedContext.outputObserver.onNext(WebSocketEvent.builder()
                .type(WebSocketSendType.SELF)
                .topic(SocketTopic.ERROR_TOPIC)
                .payload(e.getMessage())
                .build());
          } else {
            processedContext.outputObserver.onError(e);
          }
        },
        () -> playerSubjects.remove(processedContext.getUserId()));
  }
}
