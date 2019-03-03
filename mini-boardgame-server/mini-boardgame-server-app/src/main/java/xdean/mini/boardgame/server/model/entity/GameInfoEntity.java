package xdean.mini.boardgame.server.model.entity;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.OneToOne;
import javax.persistence.Table;

@Entity
@Table(name = "user_game_info")
public class GameInfoEntity {
  @Id
  long userId;

  @OneToOne
  UserEntity user;
}
