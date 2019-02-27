package xdean.wechat.mini.boardgame.server.service;

import xdean.wechat.mini.boardgame.server.model.GameBoard;
import xdean.wechat.mini.boardgame.server.model.GameRoom;

public interface GameService {
  String name();

  GameBoard createGame(GameRoom room);
}
