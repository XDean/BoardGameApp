package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Value;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

@Value
@Builder
public class CurrentGameResponse {
  @Default
  int errorCode = 0;

  GameRoomEntity room;

  @Value
  @Builder
  public static class CurrentGamePlayer {
    int id;
    int position;
  }
}
