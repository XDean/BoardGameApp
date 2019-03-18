package xdean.mini.boardgame.server.model.param;

import io.swagger.annotations.ApiModelProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class SearchGameRequest {
  @ApiModelProperty(position = 0, example = "gdjzj")
  String gameName;

  @ApiModelProperty(position = 1, example = "1")
  int page;

  @ApiModelProperty(position = 2, example = "10")
  int pageSize;
}
