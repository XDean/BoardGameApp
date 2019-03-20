package xdean.mini.boardgame.server.model;

import org.springframework.util.Assert;

import lombok.Data;
import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

@Data
public abstract class GameBoard {
  public enum State {
    OVER,
    WAITING,
    PLAYING
  }

  GameRoomEntity room;
  State state = State.WAITING;

  public abstract void start(GamePlayerEntity[] players);

  public void checkState(State state) {
    Assert.isTrue(state == this.state, "Illegal board state, expect: " + state + ", but: " + this.state);
  }
}
