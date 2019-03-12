package xdean.mini.boardgame.server.endpoint;

import java.util.Optional;

import javax.inject.Inject;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
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
import xdean.mini.boardgame.server.handler.LoginSuccessProvider;
import xdean.mini.boardgame.server.model.GlobalConstants;
import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.param.CreateGameRequest;
import xdean.mini.boardgame.server.model.param.CreateGameResponse;
import xdean.mini.boardgame.server.model.param.CurrentGameResponse;
import xdean.mini.boardgame.server.model.param.ExitGameRequest;
import xdean.mini.boardgame.server.model.param.ExitGameResponse;
import xdean.mini.boardgame.server.model.param.JoinGameRequest;
import xdean.mini.boardgame.server.model.param.JoinGameResponse;
import xdean.mini.boardgame.server.model.param.SearchGameRequest;
import xdean.mini.boardgame.server.model.param.SearchGameResponse;
import xdean.mini.boardgame.server.mybatis.mapper.GameMapper;
import xdean.mini.boardgame.server.service.GameCenterService;
import xdean.mini.boardgame.server.service.UserDataService;

@Api(tags = "Game/Game-Center")
@RestController
@RequestMapping("/game/room")
public class GameCenterEndPoint implements GlobalConstants, LoginSuccessProvider {

  private @Inject UserDataService userService;
  private @Inject GameCenterService service;
  private @Inject GameMapper gameMapper;

  @PostMapping("/create")
  @ApiOperation("Create a new game room")
  public CreateGameResponse createGame(@Valid @RequestBody CreateGameRequest request, @ApiIgnore HttpSession session) {
    CreateGameResponse response = service.createGame(request);
    if (response.getRoomId() != -1) {
      gameMapper.findRoom(response.getRoomId()).ifPresent(e -> session.setAttribute(AttrKey.ROOM, e.getRoom()));
    }
    return response;
  }

  @PostMapping("/join")
  @ApiOperation("Join an exist game room")
  public JoinGameResponse joinGame(@RequestBody JoinGameRequest request, @ApiIgnore HttpSession session) {
    JoinGameResponse response = service.joinGame(request);
    if (response.getErrorCode() == 0) {
      gameMapper.findRoom(request.getRoomId()).ifPresent(e -> session.setAttribute(AttrKey.ROOM, e.getRoom()));
    }
    return response;
  }

  @PostMapping("/exit")
  @ApiOperation("Exit game room")
  public ExitGameResponse exitGame(@RequestBody ExitGameRequest request, @ApiIgnore HttpSession session) {
    ExitGameResponse response = service.exitGame(request);
    if (response.getErrorCode() == 0) {
      session.removeAttribute(AttrKey.USER_ID);
      session.removeAttribute(AttrKey.ROOM);
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

  @Override
  public void afterSuccessLogin(HttpServletRequest request, HttpServletResponse response, String username) {
    Optional<UserEntity> user = userService.findUserByUsername(username);
    if (user.isPresent()) {
      int id = user.get().getId();
      request.getSession().setAttribute(AttrKey.USER_ID, id);
      gameMapper.findPlayer(id).ifPresent(e -> request.getSession().setAttribute(AttrKey.ROOM, e.getRoom()));
    }
  }
}
