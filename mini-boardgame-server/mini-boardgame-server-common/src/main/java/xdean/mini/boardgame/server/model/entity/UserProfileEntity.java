package xdean.mini.boardgame.server.model.entity;

import javax.persistence.Embedded;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.OneToOne;
import javax.persistence.Table;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import xdean.mini.boardgame.server.model.UserProfile;

@Data
@Entity
@NoArgsConstructor
@AllArgsConstructor
@Builder(toBuilder = true)
@Table(name = "user_profiles")
public class UserProfileEntity {
  @Id
  long userId;

  @OneToOne
  UserProfileEntity user;

  @Embedded
  UserProfile profile;
}
