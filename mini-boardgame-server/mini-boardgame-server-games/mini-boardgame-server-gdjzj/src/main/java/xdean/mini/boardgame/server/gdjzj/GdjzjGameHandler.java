package xdean.mini.boardgame.server.gdjzj;

import org.springframework.stereotype.Component;
import org.springframework.web.socket.WebSocketSession;

import com.fasterxml.jackson.databind.JsonNode;

import io.reactivex.Observable;
import xdean.mini.boardgame.server.model.GameRoom;
import xdean.mini.boardgame.server.socket.WebSocketEvent;
import xdean.mini.boardgame.server.socket.WebSocketProvider;

@Component
public class GdjzjGameHandler implements WebSocketProvider {
  @Override
  public Observable<WebSocketEvent<?>> handle(WebSocketSession session, GameRoom room,
      Observable<WebSocketEvent<JsonNode>> input) {
    return Observable.empty();
  }
}
