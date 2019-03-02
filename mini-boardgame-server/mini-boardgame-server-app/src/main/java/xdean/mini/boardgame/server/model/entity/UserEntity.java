package xdean.mini.boardgame.server.model.entity;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.OneToOne;
import javax.persistence.Table;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Entity
@AllArgsConstructor
@NoArgsConstructor
@Builder(toBuilder = true)
@Table(name = "users")
public class UserEntity {
  @Id
  long id;
  String username;
  String password;
  boolean enabled;
  @OneToOne
  @JoinColumn(name = "id")
  UserProfileEntity profile;
}
