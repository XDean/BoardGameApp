package xdean.mini.boardgame.server.model;

import org.springframework.util.Assert;

import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;

public abstract class GameBoard {
  public static final int OVER = -100;
  public static final int WAITING = -101;
  public static final int START = -102;

  protected final GameRoom room;
  protected int state = 0;

  public GameBoard(GameRoom room) {
    this.room = room;
  }

  public abstract void start(GamePlayerEntity[] players);

  public void checkState(int state) {
    Assert.isTrue(state == this.state, "Illegal board state, expect: " + state + ", but: " + this.state);
  }
}
