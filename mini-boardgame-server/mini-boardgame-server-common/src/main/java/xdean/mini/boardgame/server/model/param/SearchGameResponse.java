package xdean.mini.boardgame.server.model.param;

import java.util.List;

import lombok.Builder;
import lombok.Singular;
import lombok.Value;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

@Value
@Builder
public class SearchGameResponse {
  @Singular
  List<GameRoomEntity> rooms;
}
