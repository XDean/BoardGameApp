package xdean.mini.boardgame.server.endpoint;

import javax.inject.Inject;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import xdean.mini.boardgame.server.model.param.CreateGameRequest;
import xdean.mini.boardgame.server.model.param.CreateGameResponse;
import xdean.mini.boardgame.server.model.param.CurrentGameResponse;
import xdean.mini.boardgame.server.model.param.ExitGameRequest;
import xdean.mini.boardgame.server.model.param.ExitGameResponse;
import xdean.mini.boardgame.server.model.param.JoinGameRequest;
import xdean.mini.boardgame.server.model.param.JoinGameResponse;
import xdean.mini.boardgame.server.model.param.SearchGameRequest;
import xdean.mini.boardgame.server.model.param.SearchGameResponse;
import xdean.mini.boardgame.server.service.GameCenterService;

@RestController
@RequestMapping("/game")
public class GameCenterEndpoint {

  @Inject
  GameCenterService service;

  @GetMapping("/create")
  public CreateGameResponse createGame(CreateGameRequest request) {
    return service.createGame(request);
  }

  @GetMapping("/join")
  public JoinGameResponse joinGame(JoinGameRequest request) {
    return service.joinGame(request);
  }

  @GetMapping("/exit")
  public ExitGameResponse exitGame(ExitGameRequest request) {
    return service.exitGame(request);
  }

  @GetMapping("/search")
  SearchGameResponse searchGame(SearchGameRequest request) {
    return service.searchGame(request);
  }

  @GetMapping("")
  CurrentGameResponse currentGame() {
    return service.currentGame();
  }
}
