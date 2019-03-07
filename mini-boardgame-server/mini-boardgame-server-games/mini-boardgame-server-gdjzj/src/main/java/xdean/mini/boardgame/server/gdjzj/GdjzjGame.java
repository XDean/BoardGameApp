package xdean.mini.boardgame.server.gdjzj;

import org.springframework.stereotype.Component;
import org.springframework.web.socket.WebSocketSession;

import com.fasterxml.jackson.databind.JsonNode;

import io.reactivex.Observable;
import xdean.mini.boardgame.server.model.GameRoom;
import xdean.mini.boardgame.server.service.GameService;
import xdean.mini.boardgame.server.socket.WebSocketProvider;
import xdean.mini.boardgame.server.socket.WebSocketEvent;

@Component
public class GdjzjGame implements GameService, WebSocketProvider {

  @Override
  public String name() {
    return "gdjzj";
  }

  @Override
  public void createGame(int roomId) {

  }

  @Override
  public Observable<WebSocketEvent<?>> handle(WebSocketSession session, GameRoom room,
      Observable<WebSocketEvent<JsonNode>> input) {
    if (!room.getGameName().equals(name())) {
      return Observable.empty();
    }
    return input.map(i -> i.toBuilder().topic("gdjzj + " + i.getTopic()).build());
  }
}
