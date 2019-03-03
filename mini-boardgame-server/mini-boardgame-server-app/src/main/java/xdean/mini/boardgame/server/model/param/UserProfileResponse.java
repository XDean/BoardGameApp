package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Value;
import xdean.mini.boardgame.server.model.UserProfile;

@Value
@Builder
public class UserProfileResponse {
  public static final int USER_NOT_FOUND = 1;
  public static final int PROFILE_NOT_FOUND = 2;
  public static final int INPUT_USER = 3;

  @Default
  int errorCode = 0;

  UserProfile profile;
}
