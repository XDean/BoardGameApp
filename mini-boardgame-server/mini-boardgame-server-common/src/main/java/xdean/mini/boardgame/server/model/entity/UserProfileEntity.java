package xdean.mini.boardgame.server.model.entity;

import javax.annotation.CheckForNull;

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
  int id;

  String nickname;

  @CheckForNull
  Boolean male;

  String avatarUrl;
}
