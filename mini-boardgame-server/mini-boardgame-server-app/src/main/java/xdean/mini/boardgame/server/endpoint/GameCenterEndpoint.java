package xdean.mini.boardgame.server.endpoint;

import javax.inject.Inject;
import javax.servlet.http.HttpSession;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import xdean.mini.boardgame.server.model.GameConstants;
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
public class GameCenterEndpoint implements GameConstants {

  @Inject
  GameCenterService service;

  @PostMapping("/create")
  @ApiOperation("Create a new game room")
  public CreateGameResponse createGame(@RequestBody CreateGameRequest request, HttpSession session) {
    CreateGameResponse response = service.createGame(request);
    if (response.getRoomId() != -1) {
      session.setAttribute(ROOM_ID, response.getRoomId());
    }
    return response;
  }

  @PostMapping("/join")
  @ApiOperation("Join an exist game room")
  public JoinGameResponse joinGame(@RequestBody JoinGameRequest request, HttpSession session) {
    JoinGameResponse response = service.joinGame(request);
    if (response.getErrorCode() == 0) {
      session.setAttribute(ROOM_ID, request.getRoomId());
    }
    return response;
  }

  @PostMapping("/exit")
  @ApiOperation("Exit game room")
  public ExitGameResponse exitGame(@RequestBody ExitGameRequest request, HttpSession session) {
    ExitGameResponse response = service.exitGame(request);
    if (response.getErrorCode() == 0) {
      session.removeAttribute(ROOM_ID);
    }
    return response;
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
