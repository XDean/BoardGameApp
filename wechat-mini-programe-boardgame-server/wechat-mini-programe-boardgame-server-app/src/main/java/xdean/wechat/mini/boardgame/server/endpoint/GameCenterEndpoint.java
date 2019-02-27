package xdean.wechat.mini.boardgame.server.endpoint;

import javax.inject.Inject;

import org.springframework.web.bind.annotation.RestController;

import xdean.wechat.mini.boardgame.server.service.GameCenterService;

@RestController
public class GameCenterEndpoint {

  @Inject
  GameCenterService service;
}
