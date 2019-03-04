package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Builder.Default;
import xdean.mini.boardgame.server.model.GameRoom;
import lombok.Value;

@Value
@Builder
public class CurrentGameResponse {
  @Default
  int errorCode = 0;

  GameRoom room;
}
