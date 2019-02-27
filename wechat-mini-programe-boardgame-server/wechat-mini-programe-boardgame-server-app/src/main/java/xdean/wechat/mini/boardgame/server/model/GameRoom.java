package xdean.wechat.mini.boardgame.server.model;

import java.util.ArrayList;

import lombok.Getter;
import lombok.Setter;

public class GameRoom {
  public final int id;
  public final int playerCount;
  public final ArrayList<GamePlayer> players = new ArrayList<>();
  public @Getter @Setter GameBoard board;

  public GameRoom(int id, int playerCount) {
    this.id = id;
    this.playerCount = playerCount;
  }

  public boolean addPlayer(GamePlayer player) {
    if (players.size() < playerCount) {
      players.add(player);
      return true;
    }
    return false;
  }

  public boolean removePlayer(GamePlayer player) {
    return players.remove(player);
  }

  public boolean isEmpty() {
    return players.isEmpty();
  }
}
