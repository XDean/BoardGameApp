package xdean.mini.boardgame.server.model.param;

import javax.validation.constraints.NotNull;

import lombok.Builder;
import lombok.Value;
import xdean.mini.boardgame.server.model.UserProfile;

@Value
@Builder
public class UserProfileUpdateRequest {
  @NotNull
  UserProfile profile;
}
