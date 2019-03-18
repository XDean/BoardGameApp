package xdean.mini.boardgame.server.model.param;

import io.swagger.annotations.ApiModelProperty;
import lombok.Builder;
import lombok.Value;
import lombok.Builder.Default;

@Value
@Builder
public class CreateGameResponse {
  @Default
  @ApiModelProperty(example = "970711")
  int roomId = -1;
}
