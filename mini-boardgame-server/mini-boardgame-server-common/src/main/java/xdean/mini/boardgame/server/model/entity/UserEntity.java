package xdean.mini.boardgame.server.model.entity;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.Builder.Default;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder(toBuilder = true)
public class UserEntity {
  @Default
  int id = -1;
  String username;
  String password;
  boolean enabled;

  List<String> authorities;

  @Default
  UserProfileEntity profile = UserProfileEntity.builder().build();
}
