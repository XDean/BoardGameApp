package xdean.mini.boardgame.server.gdjzj;

import org.springframework.stereotype.Component;
import org.springframework.util.Assert;

import xdean.mini.boardgame.server.gdjzj.game.GdjzjBoard;
import xdean.mini.boardgame.server.gdjzj.model.GdjzjConfig;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.service.GameProvider;

@Component
public class GdjzjGame implements GameProvider<GdjzjBoard, GdjzjConfig> {

  @Override
  public String name() {
    return "gdjzj";
  }

  @Override
  public void configRoom(GameRoomEntity room, GdjzjConfig config) {
    GameProvider.super.configRoom(room, config);
    room.setPlayerCount(config.getPlayerCount());
  }

  @Override
  public GdjzjBoard createGame(GameRoomEntity room) {
    Assert.isTrue(room.getGameName().equals(name()), "Game must be gdjzj");
    Assert.isTrue(room.getPlayerCount() >= 6 && room.getPlayerCount() <= 8, "Gdjzj player count must be 6-8");
    return new GdjzjBoard();
  }

  @Override
  public Class<GdjzjConfig> configClass() {
    return GdjzjConfig.class;
  }
}
