package xdean.wechat.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Value;

@Value
@Builder
public class ExitGameRequest {
  int roomId;
}