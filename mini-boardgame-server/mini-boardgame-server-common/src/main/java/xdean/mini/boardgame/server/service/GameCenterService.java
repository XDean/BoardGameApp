package xdean.mini.boardgame.server.service;

import xdean.mini.boardgame.server.model.GameRoom;
import xdean.mini.boardgame.server.model.param.CreateGameRequest;
import xdean.mini.boardgame.server.model.param.CreateGameResponse;
import xdean.mini.boardgame.server.model.param.CurrentGameResponse;
import xdean.mini.boardgame.server.model.param.ExitGameRequest;
import xdean.mini.boardgame.server.model.param.ExitGameResponse;
import xdean.mini.boardgame.server.model.param.JoinGameRequest;
import xdean.mini.boardgame.server.model.param.JoinGameResponse;
import xdean.mini.boardgame.server.model.param.SearchGameRequest;
import xdean.mini.boardgame.server.model.param.SearchGameResponse;

public interface GameCenterService {
  CreateGameResponse createGame(CreateGameRequest request);

  JoinGameResponse joinGame(JoinGameRequest request);

  ExitGameResponse exitGame(ExitGameRequest request);

  SearchGameResponse searchGame(SearchGameRequest request);

  CurrentGameResponse currentGame();
}