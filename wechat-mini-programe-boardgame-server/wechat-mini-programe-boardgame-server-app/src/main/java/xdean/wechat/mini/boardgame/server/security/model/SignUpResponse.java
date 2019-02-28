package xdean.wechat.mini.boardgame.server.security.model;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Value;

@Value
@Builder
public class SignUpResponse {
  public static final int INPUT_USERNAME_PASSWORD = 1;
  public static final int ILLEGAL_USERNAME = 2;
  public static final int ILLEGAL_PASSWORD = 3;
  public static final int USERNAME_EXIST = 4;

  boolean success;

  @Default
  int errorCode = -1;

  @Default
  String message = "";
}
