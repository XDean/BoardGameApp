package xdean.wechat.mini.boardgame.server.service;

import xdean.wechat.mini.boardgame.server.model.param.CreateGameRequest;
import xdean.wechat.mini.boardgame.server.model.param.CreateGameResponse;
import xdean.wechat.mini.boardgame.server.model.param.JoinGameRequest;
import xdean.wechat.mini.boardgame.server.model.param.JoinGameResponse;

public interface GameCenterService {
  CreateGameResponse createGame(CreateGameRequest request);

  JoinGameResponse joinGame(JoinGameRequest request);
}
