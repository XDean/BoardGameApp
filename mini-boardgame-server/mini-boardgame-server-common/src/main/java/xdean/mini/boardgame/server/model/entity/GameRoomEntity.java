package xdean.mini.boardgame.server.model.entity;

import java.util.List;

import javax.persistence.AttributeOverride;
import javax.persistence.Column;
import javax.persistence.Embedded;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.OneToMany;
import javax.persistence.Table;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.Singular;
import xdean.mini.boardgame.server.model.GameRoom;

@Data
@Entity
@AllArgsConstructor
@NoArgsConstructor
@Builder(toBuilder = true)
@Table(name = "game_rooms")
public class GameRoomEntity {
  @Id
  int id;

  @Embedded
  @AttributeOverride(name = "id", column = @Column(insertable = false, updatable = false))
  GameRoom room;

  @Singular
  @OneToMany(mappedBy = "room")
  List<GamePlayerEntity> players;

  public void setPlayers(List<GamePlayerEntity> players) {
    this.players = players;
    if (room != null) {
      room.setCurrentPlayerCount(players.size());
    }
  }

  public void setRoom(GameRoom room) {
    this.room = room;
    if (players != null) {
      room.setCurrentPlayerCount(players.size());
    }
  }

  public void addPlayer(GamePlayerEntity player) {
    players.add(player);
    room.setCurrentPlayerCount(players.size());
  }

  public void removePlayer(GamePlayerEntity player) {
    players.remove(player);
    room.setCurrentPlayerCount(players.size());
  }
}
