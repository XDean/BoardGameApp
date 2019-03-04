package xdean.mini.boardgame.server.model.param;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class SearchGameRequest {
  String gameName;
  int limit;
  int offset;
}
