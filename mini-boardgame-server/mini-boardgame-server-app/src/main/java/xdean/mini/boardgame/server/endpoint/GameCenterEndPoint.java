package xdean.mini.boardgame.server.endpoint;

import javax.inject.Inject;
import javax.servlet.http.HttpSession;
import javax.validation.Valid;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import springfox.documentation.annotations.ApiIgnore;
import xdean.mini.boardgame.server.model.CommonConstants;
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
import xdean.mini.boardgame.server.service.GameDataService;

@Api(tags = "Game/Game-Center")
@RestController
@RequestMapping("/game/room")
public class GameCenterEndPoint implements CommonConstants {

  private @Inject GameCenterService service;
  private @Inject GameDataService gameService;

  @PostMapping("/create")
  @ApiOperation("Create a new game room")
  public CreateGameResponse createGame(@Valid @RequestBody CreateGameRequest request, @ApiIgnore HttpSession session) {
    CreateGameResponse response = service.createGame(request);
    if (response.getRoomId() != -1) {
      gameService.findRoom(response.getRoomId()).ifPresent(e -> session.setAttribute(AttrKey.ROOM, e));
    }
    return response;
  }

  @PostMapping("/join")
  @ApiOperation("Join an exist game room")
  public JoinGameResponse joinGame(@RequestBody JoinGameRequest request, @ApiIgnore HttpSession session) {
    JoinGameResponse response = service.joinGame(request);
    gameService.findRoom(request.getRoomId()).ifPresent(e -> session.setAttribute(AttrKey.ROOM, e));
    return response;
  }

  @PostMapping("/exit")
  @ApiOperation("Exit game room")
  public ExitGameResponse exitGame(@RequestBody ExitGameRequest request, @ApiIgnore HttpSession session) {
    ExitGameResponse response = service.exitGame(request);
    session.removeAttribute(AttrKey.USER_ID);
    session.removeAttribute(AttrKey.ROOM);
    return response;
  }

  @GetMapping("/search")
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
