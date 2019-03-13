package xdean.mini.boardgame.server.model.param;

import java.util.List;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Singular;
import lombok.Value;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

@Value
@Builder
public class SearchGameResponse {
  @Default
  int errorCode = 0;

  @Singular
  List<GameRoomEntity> rooms;
}
