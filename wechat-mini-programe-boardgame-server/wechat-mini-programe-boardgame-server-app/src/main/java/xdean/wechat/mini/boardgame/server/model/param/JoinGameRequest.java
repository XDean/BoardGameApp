package xdean.wechat.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Value;

@Value
@Builder
public class JoinGameRequest {
  String name;
  int roomId;
}
