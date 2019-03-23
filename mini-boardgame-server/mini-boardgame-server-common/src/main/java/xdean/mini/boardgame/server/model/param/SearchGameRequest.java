package xdean.mini.boardgame.server.model.param;

import io.swagger.annotations.ApiModelProperty;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
public class SearchGameRequest {
  @ApiModelProperty(position = 0, required = false, example = "gdjzj")
  String gameName = "";

  @ApiModelProperty(position = 1, required = false, example = "-1")
  int roomId = -1;

  @ApiModelProperty(position = 2, required = false, example = "0")
  int page = 0;

  @ApiModelProperty(position = 3, required = false, example = "20")
  int pageSize = 20;
}
