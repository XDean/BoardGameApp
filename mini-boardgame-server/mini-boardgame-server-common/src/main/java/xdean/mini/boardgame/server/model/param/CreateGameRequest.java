package xdean.mini.boardgame.server.model.param;

import javax.validation.constraints.Min;
import javax.validation.constraints.NotBlank;

import io.swagger.annotations.ApiModelProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Builder.Default;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class CreateGameRequest {
  @Default
  @NotBlank
  @ApiModelProperty(position = 0, required = true, example = "gdjzj")
  String gameName = "";

  @Default
  @ApiModelProperty(position = 1, example = "New Room")
  String roomName = "";

  @Min(1)
  @ApiModelProperty(position = 2, required = true, example = "6")
  int playerCount;
}
