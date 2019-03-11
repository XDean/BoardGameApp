package xdean.mini.boardgame.server.model.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import xdean.mini.boardgame.server.model.UserProfile;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder(toBuilder = true)
public class UserProfileEntity {

  int userId;

  UserEntity user;

  UserProfile profile;
}
