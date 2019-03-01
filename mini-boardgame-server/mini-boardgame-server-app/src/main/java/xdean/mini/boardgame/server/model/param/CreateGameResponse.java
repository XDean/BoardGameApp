package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Value;
import lombok.Builder.Default;

@Value
@Builder
public class CreateGameResponse {
  boolean success;

  @Default
  int roomId = -1;

  @Default
  int errorCode = -1;
}
