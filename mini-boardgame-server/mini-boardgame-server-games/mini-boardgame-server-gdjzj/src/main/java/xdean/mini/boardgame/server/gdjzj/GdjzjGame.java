package xdean.mini.boardgame.server.gdjzj;

import org.springframework.stereotype.Component;
import org.springframework.util.Assert;
import org.springframework.web.socket.WebSocketSession;

import com.fasterxml.jackson.databind.JsonNode;

import io.reactivex.Observable;
import xdean.mini.boardgame.server.gdjzj.game.GdjzjBoard;
import xdean.mini.boardgame.server.model.GameRoom;
import xdean.mini.boardgame.server.service.GameService;
import xdean.mini.boardgame.server.socket.WebSocketEvent;
import xdean.mini.boardgame.server.socket.WebSocketProvider;

@Component
public class GdjzjGame implements GameService<GdjzjBoard>, WebSocketProvider {

  @Override
  public String name() {
    return "gdjzj";
  }

  @Override
  public GdjzjBoard createGame(GameRoom room) {
    Assert.isTrue(room.getGameName().equals(name()), "Game must be gdjzj");
    Assert.isTrue(room.getPlayerCount() >= 6 && room.getPlayerCount() <= 8, "Gdjzj player count must be 6-8");
    return new GdjzjBoard(room);
  }

  @Override
  public Observable<WebSocketEvent<?>> handle(WebSocketSession session, GameRoom room,
      Observable<WebSocketEvent<JsonNode>> input) {
    return Observable.empty();
  }
}
