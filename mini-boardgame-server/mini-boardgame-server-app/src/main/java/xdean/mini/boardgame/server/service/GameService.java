package xdean.mini.boardgame.server.service;

import xdean.mini.boardgame.server.model.GameBoard;
import xdean.mini.boardgame.server.model.GameRoom;

public interface GameService {
  String name();

  GameBoard createGame(GameRoom room);
}
