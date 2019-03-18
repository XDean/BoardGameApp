package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Value;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

@Value
@Builder
public class CurrentGameResponse {
  GameRoomEntity room;
}
