package xdean.wechat.mini.boardgame.server.security.model;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Value;

@Value
@Builder
public class LoginResponse {
  public static final int INPUT_USERNAME_PASSWORD = 1;
  public static final int WRONG_PASSWORD_OR_USERNAME_NOT_EXIST = 2;

  boolean success;

  @Default
  int errorCode = -1;

  @Default
  String message = "";
}
