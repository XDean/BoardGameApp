package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Value;

@Value
@Builder
public class SimpleResponse {
  @Default
  int errorCode = 0;

  @Default
  String errorMessage = "";
}
