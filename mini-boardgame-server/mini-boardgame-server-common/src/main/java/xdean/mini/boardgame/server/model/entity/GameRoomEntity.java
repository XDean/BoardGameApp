package xdean.mini.boardgame.server.model.entity;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnore;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.Singular;
import xdean.mini.boardgame.server.model.GameBoard;

@Data
@AllArgsConstructor
@Builder(toBuilder = true)
public class GameRoomEntity {
  int id;

  String gameName;

  int playerCount;

  String roomName;

  Date createdTime;

  @Singular
  List<GamePlayerEntity> players;

  @JsonIgnore
  GameBoard board;

  public GameRoomEntity() {
    players = new ArrayList<>();
  }

  public void setPlayers(List<GamePlayerEntity> players) {
    this.players.clear();
    this.players.addAll(players);
  }

  public void addPlayer(GamePlayerEntity player) {
    players.add(player);
  }

  public void removePlayer(GamePlayerEntity player) {
    players.remove(player);
  }
}
