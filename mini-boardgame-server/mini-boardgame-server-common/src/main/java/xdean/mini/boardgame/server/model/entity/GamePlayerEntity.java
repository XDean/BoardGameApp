package xdean.mini.boardgame.server.model.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class GamePlayerEntity {
  int userId;

  GameRoomEntity room;

  int seat;
}
