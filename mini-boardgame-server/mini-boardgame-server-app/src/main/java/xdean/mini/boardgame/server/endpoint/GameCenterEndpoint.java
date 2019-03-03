package xdean.mini.boardgame.server.endpoint;

import javax.inject.Inject;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import xdean.mini.boardgame.server.service.GameCenterService;

@RestController("/game")
public class GameCenterEndpoint {

  @Inject
  GameCenterService service;

  @GetMapping("/create")
  public void createGame() {

  }

  @GetMapping("/join")
  public void joinGame() {

  }

  @GetMapping("/exit")
  public void exitGame() {

  }
}
