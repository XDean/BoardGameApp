package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Data;
import xdean.mini.boardgame.server.model.UserProfile;

@Data
@Builder
public class UserProfileResponse {
  public static final int USER_NOT_FOUND = 1;
  public static final int INPUT_USER = 2;

  @Default
  int errorCode = 0;

  UserProfile profile;
}
