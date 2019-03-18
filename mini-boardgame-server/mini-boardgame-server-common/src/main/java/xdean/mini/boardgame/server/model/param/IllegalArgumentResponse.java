package xdean.mini.boardgame.server.model.param;

import lombok.AllArgsConstructor;
import lombok.Value;
import xdean.mini.boardgame.server.model.GlobalConstants;

@Value
@AllArgsConstructor
public class IllegalArgumentResponse {
  final int errorCode = GlobalConstants.ErrorCode.ILLEGAL_ARGUEMENT;
  String errorMessage;
}
