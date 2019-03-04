package xdean.mini.boardgame.server.model.param;

import java.util.List;

import lombok.Builder;
import lombok.Value;
import xdean.mini.boardgame.server.model.GameRoom;

@Value
@Builder
public class SearchGameResponse {
  List<GameRoom> rooms;
}
