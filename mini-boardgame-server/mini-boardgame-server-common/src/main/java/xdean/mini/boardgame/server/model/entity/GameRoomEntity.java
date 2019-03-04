package xdean.mini.boardgame.server.model.entity;

import java.util.List;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.OneToMany;
import javax.persistence.Table;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.Singular;

@Data
@Entity
@AllArgsConstructor
@NoArgsConstructor
@Builder(toBuilder = true)
@Table(name = "game_rooms")
public class GameRoomEntity {
  @Id
  int id;

  String gameName;

  @Singular
  @OneToMany(mappedBy = "room")
  List<GamePlayerEntity> players;
}
