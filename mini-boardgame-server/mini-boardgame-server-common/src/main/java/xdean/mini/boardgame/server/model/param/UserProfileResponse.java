package xdean.mini.boardgame.server.model.param;

import lombok.Builder;
import lombok.Value;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;

@Value
@Builder
public class UserProfileResponse {

  UserProfileEntity profile;
}
