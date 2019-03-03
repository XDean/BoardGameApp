package xdean.mini.boardgame.server.model.param;

import javax.validation.constraints.NotEmpty;

import lombok.Builder;
import lombok.Value;

@Value
@Builder
public class JoinGameRequest {
  @NotEmpty
  String gameName;

  int roomId;
}
