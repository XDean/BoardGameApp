package xdean.mini.boardgame.server.gdjzj;

import org.springframework.stereotype.Component;
import org.springframework.util.Assert;

import xdean.mini.boardgame.server.gdjzj.game.GdjzjBoard;
import xdean.mini.boardgame.server.model.GameRoom;
import xdean.mini.boardgame.server.service.GameService;

@Component
public class GdjzjGame implements GameService<GdjzjBoard> {

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
}
