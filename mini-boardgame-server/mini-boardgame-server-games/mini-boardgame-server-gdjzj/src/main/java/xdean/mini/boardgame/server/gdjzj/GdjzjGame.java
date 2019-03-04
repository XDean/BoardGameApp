package xdean.mini.boardgame.server.gdjzj;

import org.springframework.stereotype.Component;

import xdean.mini.boardgame.server.service.GameService;

@Component
public class GdjzjGame implements GameService {

  @Override
  public String name() {
    return "gdjzj";
  }

  @Override
  public void createGame(int roomId) {

  }
}
