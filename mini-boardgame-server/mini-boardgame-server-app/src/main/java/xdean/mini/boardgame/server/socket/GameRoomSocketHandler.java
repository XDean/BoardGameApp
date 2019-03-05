package xdean.mini.boardgame.server.socket;

import java.util.List;
import java.util.concurrent.CopyOnWriteArrayList;

import org.springframework.web.socket.WebSocketSession;

public class GameRoomSocketHandler {
  private final int roomId;
  final List<WebSocketSession> sessions = new CopyOnWriteArrayList<>();

  public GameRoomSocketHandler(int roomId) {
    this.roomId = roomId;
  }

  public void handle(String message) {

  }
}
