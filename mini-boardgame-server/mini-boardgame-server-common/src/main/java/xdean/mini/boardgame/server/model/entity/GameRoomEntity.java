package xdean.mini.boardgame.server.model.entity;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnore;

import io.swagger.annotations.ApiModelProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.Singular;
import xdean.mini.boardgame.server.model.GameBoard;

@Data
@AllArgsConstructor
@Builder(toBuilder = true)
public class GameRoomEntity {
  @ApiModelProperty(position = 0, example = "950317")
  int id;

  @ApiModelProperty(position = 1, example = "gdjzj")
  String gameName;

  @ApiModelProperty(position = 2, example = "6")
  int playerCount;

  @ApiModelProperty(position = 3, example = "New Room")
  String roomName;

  @ApiModelProperty(position = 4)
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
