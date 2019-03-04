package xdean.mini.boardgame.server.model.param;

import javax.validation.constraints.Min;

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
  String gameName = "";

  @Default
  String roomName = "";

  @Min(1)
  int playerCount;
}
