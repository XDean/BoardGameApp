package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Data;
import xdean.mini.boardgame.server.model.UserProfile;

@Data
@Builder
public class UserProfileUpdateResponse {
  public static final int NOT_CURRENT_USER = 1;
  public static final int ILLEGAL_INPUT = 2;

  @Default
  int errorCode = 0;

  UserProfile profile;
}
