package xdean.mini.boardgame.server.socket.handler;

import org.springframework.stereotype.Component;

import com.fasterxml.jackson.databind.JsonNode;

import io.reactivex.Observable;
import lombok.extern.slf4j.Slf4j;
import xdean.mini.boardgame.server.model.CommonConstants.AttrKey;
import xdean.mini.boardgame.server.model.CommonConstants.SocketTopic;
import xdean.mini.boardgame.server.socket.AbstractGameSocketProvider;
import xdean.mini.boardgame.server.socket.WebSocketEvent;
import xdean.mini.boardgame.server.socket.WebSocketSendType;

@Slf4j
@Component
public class HeartBeatHandler extends AbstractGameSocketProvider {
  @Override
  protected Observable<WebSocketEvent<JsonNode>> processInputFlow(SocketContext context) {
    return context.inputFlow
        .doOnNext(e -> {
          if (e.getTopic().equals(SocketTopic.HEART_BEAT)) {
            e.setConsumed(true);
            long sendTime = ((Number) e.getAttributes().get(AttrKey.TIMESTAMP)).longValue();
            long time = System.currentTimeMillis();
            if (time - sendTime > 5000) {
              log.info("Heart beat very delay: " + context.session);
            }
            sendEvent(context.userId, WebSocketEvent.builder()
                .id(e.getId())
                .type(WebSocketSendType.SELF)
                .topic(SocketTopic.HEART_BEAT)
                .attribute(AttrKey.TIMESTAMP, time)
                .build());
          }
        });
  }
}
