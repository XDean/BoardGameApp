package xdean.mini.boardgame.server.socket;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import org.springframework.stereotype.Component;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketSession;
import org.springframework.web.socket.handler.TextWebSocketHandler;

import xdean.jex.log.Logable;
import xdean.mini.boardgame.server.model.GameConstants;

@Component
public class GameSocketHandler extends TextWebSocketHandler implements Logable, GameConstants {

  private final Map<Integer, GameRoomSocketHandler> rooms = new ConcurrentHashMap<>();

  @Override
  public void afterConnectionEstablished(WebSocketSession session) throws Exception {
    Integer id = (Integer) session.getAttributes().get(ROOM_ID);
    if (id != null) {
      trace(session.getRemoteAddress() + " connect to room " + id);
      GameRoomSocketHandler room = rooms.computeIfAbsent(id, GameRoomSocketHandler::new);
      room.sessions.add(session);
    }
  }

  @Override
  public void afterConnectionClosed(WebSocketSession session, CloseStatus status) throws Exception {
    Integer id = (Integer) session.getAttributes().get(ROOM_ID);
    if (id != null) {
      GameRoomSocketHandler room = rooms.get(id);
      if (room != null) {
        room.sessions.remove(session);
      }
    }
  }

  @Override
  protected void handleTextMessage(WebSocketSession session, TextMessage message) throws Exception {
    Integer id = (Integer) session.getAttributes().get(ROOM_ID);
    if (id != null) {
      GameRoomSocketHandler room = rooms.get(id);
      if (room != null) {
        room.handle(message.getPayload());
      }
    }
  }
}
