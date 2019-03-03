package xdean.mini.boardgame.server.model.param;

import javax.validation.constraints.Min;
import javax.validation.constraints.NotEmpty;

import lombok.Builder;
import lombok.Value;

@Value
@Builder
public class CreateGameRequest {
  @NotEmpty
  String gameName;

  @Min(1)
  int playerCount;
}
