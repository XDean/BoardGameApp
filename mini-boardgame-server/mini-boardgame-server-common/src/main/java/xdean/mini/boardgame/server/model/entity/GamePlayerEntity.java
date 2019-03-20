package xdean.mini.boardgame.server.model.entity;

import java.util.Optional;

import com.fasterxml.jackson.annotation.JsonIgnore;

import io.swagger.annotations.ApiModelProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.Builder.Default;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class GamePlayerEntity {
  @ApiModelProperty(position = 0, example = "1021")
  int id;

  @Default
  @JsonIgnore
  Optional<GameRoomEntity> room = Optional.empty();

  @Default
  @ApiModelProperty(position = 1, example = "0")
  int seat = -1;

  boolean ready;

  public void setRoom(GameRoomEntity e) {
    room = Optional.ofNullable(e);
  }
}
