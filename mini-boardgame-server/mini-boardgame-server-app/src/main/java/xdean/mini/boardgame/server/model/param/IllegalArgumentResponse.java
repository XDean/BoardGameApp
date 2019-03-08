package xdean.mini.boardgame.server.model.param;

import lombok.AllArgsConstructor;
import lombok.Value;
import xdean.mini.boardgame.server.model.GameConstants;

@Value
@AllArgsConstructor
public class IllegalArgumentResponse {
  final int errorCode = GameConstants.ErrorCode.ILLEGAL_ARGUEMENT;
  String errorMessage;
}
