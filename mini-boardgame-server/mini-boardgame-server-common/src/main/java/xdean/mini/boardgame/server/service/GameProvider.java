package xdean.mini.boardgame.server.service;

import xdean.mini.boardgame.server.model.GameBoard;
import xdean.mini.boardgame.server.model.GameConfig;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

public interface GameProvider<T extends GameBoard, C extends GameConfig> {
  String name();

  default void configRoom(GameRoomEntity room, C config) {
    room.setConfig(config);
  }

  T createGame(GameRoomEntity room) throws IllegalArgumentException;

  Class<C> configClass();
}
