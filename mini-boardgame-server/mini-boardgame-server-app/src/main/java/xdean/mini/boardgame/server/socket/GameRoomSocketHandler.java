package xdean.mini.boardgame.server.socket;

import java.util.List;
import java.util.concurrent.CopyOnWriteArrayList;

import org.springframework.web.socket.WebSocketSession;

import xdean.mini.boardgame.server.model.GameRoom;

public class GameRoomSocketHandler {
  private final GameRoom room;
  final List<WebSocketSession> sessions = new CopyOnWriteArrayList<>();

  public GameRoomSocketHandler(GameRoom room) {
    this.room = room;
  }

  public void handleMessage(String message) {

  }

  public void sendMessage() {

  }
}
