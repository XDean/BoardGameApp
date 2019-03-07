package xdean.mini.boardgame.server.model;

import java.util.Date;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class GameRoom {

  int id;

  String gameName;

  int playerCount;

  String roomName;

  Date createdTime;
}
