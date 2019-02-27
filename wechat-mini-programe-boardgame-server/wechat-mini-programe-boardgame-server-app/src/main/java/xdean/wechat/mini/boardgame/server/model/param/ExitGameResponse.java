package xdean.wechat.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Value;
import lombok.Builder.Default;

@Value
@Builder
public class ExitGameResponse {
  boolean success;

  @Default
  int errorCode = -1;
}
