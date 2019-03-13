package xdean.mini.boardgame.server.model.entity;

import java.util.Date;
import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.Singular;
import xdean.mini.boardgame.server.model.GameBoard;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder(toBuilder = true)
public class GameRoomEntity {
  int id;

  String gameName;

  int playerCount;

  String roomName;

  Date createdTime;

  @Singular
  List<GamePlayerEntity> players;

  GameBoard board;

  public void addPlayer(GamePlayerEntity player) {
    players.add(player);
  }

  public void removePlayer(GamePlayerEntity player) {
    players.remove(player);
  }
}
