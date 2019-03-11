package xdean.mini.boardgame.server.service;

import xdean.mini.boardgame.server.model.GameBoard;
import xdean.mini.boardgame.server.model.GameRoom;

public interface GameProvider<T extends GameBoard> {
  String name();

  T createGame(GameRoom room) throws IllegalArgumentException;
}
