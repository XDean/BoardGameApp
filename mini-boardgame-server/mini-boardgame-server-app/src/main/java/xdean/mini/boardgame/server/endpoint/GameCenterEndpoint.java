package xdean.mini.boardgame.server.endpoint;

import javax.inject.Inject;

import org.springframework.web.bind.annotation.RestController;

import xdean.mini.boardgame.server.service.GameCenterService;

@RestController
public class GameCenterEndpoint {

  @Inject
  GameCenterService service;
}
