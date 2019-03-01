package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Value;
import lombok.Builder.Default;

@Value
@Builder
public class JoinGameResponse {
  boolean success;

  @Default
  int errorCode = -1;
}
