package xdean.mini.boardgame.server.model.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder(toBuilder = true)
public class UserEntity {
  int id;
  String username;
  String password;
  boolean enabled;

  UserProfileEntity profile;
}
