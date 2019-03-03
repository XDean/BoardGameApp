package xdean.mini.boardgame.server.model;

import lombok.Getter;
import lombok.Setter;

public class GamePlayer {
  public @Getter final int id;
  private @Getter @Setter int roomId = -1;

  public GamePlayer(int id) {
    this.id = id;
  }
}
