package xdean.mini.boardgame.server.security.model;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Value;

@Value
@Builder
public class LoginOpenIdResponse {
  public static final int BAD_CREDENTIALS = 1;
  public static final int PROVIDER_NOT_FOUND = 2;
  public static final int PROVIDE_TOKEN_PROVIDER = 3;

  boolean success;

  @Default
  int errorCode = 0;

  @Default
  String message = "";
}
