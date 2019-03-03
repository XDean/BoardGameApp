package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Data;

@Data
@Builder
public class JoinGameResponse {
  @Default
  int errorCode = 0;
}
