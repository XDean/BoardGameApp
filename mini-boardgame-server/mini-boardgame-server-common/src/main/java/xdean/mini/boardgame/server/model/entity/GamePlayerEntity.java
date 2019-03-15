package xdean.mini.boardgame.server.model.entity;

import java.util.Optional;

import com.fasterxml.jackson.annotation.JsonIgnore;

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
  int id;

  @Default
  @JsonIgnore
  Optional<GameRoomEntity> room = Optional.empty();

  @Default
  int seat = -1;

  public void setRoom(GameRoomEntity e) {
    room = Optional.ofNullable(e);
  }
}
