package xdean.mini.boardgame.server.model.param;

import io.swagger.annotations.ApiModelProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.Builder.Default;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class SearchGameRequest {
  @ApiModelProperty(position = 0, example = "gdjzj")
  String gameName;

  @Default
  @ApiModelProperty(position = 1, example = "0")
  int page = 0;

  @Default
  @ApiModelProperty(position = 2, example = "20")
  int pageSize = 20;
}
