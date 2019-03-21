package xdean.mini.boardgame.server.model.param;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;

import com.fasterxml.jackson.databind.JsonNode;

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

  @NotNull
  @ApiModelProperty(position = 2, required = true)
  JsonNode gameConfig;
}
