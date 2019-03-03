package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Value;
import xdean.mini.boardgame.server.model.UserProfile;

@Value
@Builder
public class UserProfileUpdateResponse {
  public static final int HAVE_NOT_LOGIN = 1;
  public static final int ILLEGAL_INPUT = 2;

  @Default
  int errorCode = 0;

  UserProfile profile;
}
