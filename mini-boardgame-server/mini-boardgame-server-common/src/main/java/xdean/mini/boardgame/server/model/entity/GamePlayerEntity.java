package xdean.mini.boardgame.server.model.entity;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Entity
@Builder
@NoArgsConstructor
@AllArgsConstructor
@Table(name = "game_players")
public class GamePlayerEntity {
  @Id
  int userId;

  @ManyToOne
  GameRoomEntity room;

  int seat;
}
