package xdean.wechat.mini.boardgame.server.service;

import xdean.wechat.mini.boardgame.server.model.GamePlayer;
import xdean.wechat.mini.boardgame.server.model.param.CreateGameRequest;
import xdean.wechat.mini.boardgame.server.model.param.CreateGameResponse;
import xdean.wechat.mini.boardgame.server.model.param.ExitGameRequest;
import xdean.wechat.mini.boardgame.server.model.param.ExitGameResponse;
import xdean.wechat.mini.boardgame.server.model.param.JoinGameRequest;
import xdean.wechat.mini.boardgame.server.model.param.JoinGameResponse;

public interface GameCenterService {
  CreateGameResponse createGame(GamePlayer player, CreateGameRequest request);

  JoinGameResponse joinGame(GamePlayer player, JoinGameRequest request);

  ExitGameResponse exitGame(GamePlayer player, ExitGameRequest request);
}
