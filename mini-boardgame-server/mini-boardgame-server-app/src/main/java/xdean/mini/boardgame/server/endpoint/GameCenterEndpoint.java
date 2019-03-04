package xdean.mini.boardgame.server.endpoint;

import javax.inject.Inject;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
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

@Api(tags = "Game/Game-Center")
@RestController
@RequestMapping("/game")
public class GameCenterEndpoint {

  @Inject
  GameCenterService service;

  @PostMapping("/create")
  @ApiOperation("Create a new game room")
  public CreateGameResponse createGame(@RequestBody CreateGameRequest request) {
    return service.createGame(request);
  }

  @PostMapping("/join")
  @ApiOperation("Join an exist game room")
  public JoinGameResponse joinGame(@RequestBody JoinGameRequest request) {
    return service.joinGame(request);
  }

  @PostMapping("/exit")
  @ApiOperation("Exit game room")
  public ExitGameResponse exitGame(@RequestBody ExitGameRequest request) {
    return service.exitGame(request);
  }

  @PostMapping("/search")
  @ApiOperation("Search exist game rooms")
  SearchGameResponse searchGame(@RequestBody SearchGameRequest request) {
    return service.searchGame(request);
  }

  @GetMapping("")
  @ApiOperation("Get current user room")
  CurrentGameResponse currentGame() {
    return service.currentGame();
  }
}
