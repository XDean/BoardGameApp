package xdean.mini.boardgame.server.service;

import xdean.mini.boardgame.server.model.GameBoard;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

public interface GameProvider<T extends GameBoard> {
  String name();

  T createGame(GameRoomEntity room) throws IllegalArgumentException;
}
