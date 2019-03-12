package xdean.mini.boardgame.server.model.entity;

import com.fasterxml.jackson.annotation.JsonIgnore;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder(toBuilder = true)
public class UserProfileEntity {

  @JsonIgnore
  int userId;

  String nickname;
  boolean male;
  String avatarUrl;
}
