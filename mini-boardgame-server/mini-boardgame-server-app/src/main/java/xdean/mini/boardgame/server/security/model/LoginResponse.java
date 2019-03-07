package xdean.mini.boardgame.server.security.model;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Value;

@Value
@Builder
public class LoginResponse {
  public static final int BAD_INPUT = 1;
  public static final int BAD_CREDENTIALS = 2;
  public static final int PROVIDER_NOT_FOUND = 3;

  @Default
  int errorCode = 0;

  @Default
  String message = "";
}
