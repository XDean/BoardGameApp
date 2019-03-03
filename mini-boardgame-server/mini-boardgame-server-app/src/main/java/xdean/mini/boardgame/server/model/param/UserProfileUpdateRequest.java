package xdean.mini.boardgame.server.model.param;

import javax.validation.constraints.NotNull;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import xdean.mini.boardgame.server.model.UserProfile;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class UserProfileUpdateRequest {
  @NotNull
  UserProfile profile;
}
