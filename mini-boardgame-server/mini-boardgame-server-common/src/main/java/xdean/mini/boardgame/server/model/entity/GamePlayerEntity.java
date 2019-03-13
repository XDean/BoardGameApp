package xdean.mini.boardgame.server.model.entity;

import com.fasterxml.jackson.annotation.JsonIgnore;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class GamePlayerEntity {
  int id;

  @JsonIgnore
  GameRoomEntity room;

  int seat;
}
