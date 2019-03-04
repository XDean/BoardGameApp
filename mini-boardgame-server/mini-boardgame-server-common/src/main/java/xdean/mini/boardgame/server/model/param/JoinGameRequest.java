package xdean.mini.boardgame.server.model.param;

import javax.validation.constraints.NotEmpty;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class JoinGameRequest {
  @NotEmpty
  String gameName;

  int roomId;
}
