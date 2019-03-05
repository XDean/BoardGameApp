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
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

@Component
public class GameSocketHandler extends TextWebSocketHandler implements Logable, GameConstants {

  private final Map<Integer, GameRoomSocketHandler> rooms = new ConcurrentHashMap<>();

  @Override
  public void afterConnectionEstablished(WebSocketSession session) throws Exception {
    GameRoomEntity e = (GameRoomEntity) session.getAttributes().get(ROOM);
    if (e != null) {
      trace(session.getRemoteAddress() + " connect to room " + e.getId());
      GameRoomSocketHandler room = rooms.computeIfAbsent(e.getId(), i -> new GameRoomSocketHandler(e.getRoom()));
      room.sessions.add(session);
    }
  }

  @Override
  public void afterConnectionClosed(WebSocketSession session, CloseStatus status) throws Exception {
    GameRoomEntity e = (GameRoomEntity) session.getAttributes().get(ROOM);
    if (e != null) {
      GameRoomSocketHandler room = rooms.get(e.getId());
      if (room != null) {
        room.sessions.remove(session);
      }
    }
  }

  @Override
  protected void handleTextMessage(WebSocketSession session, TextMessage message) throws Exception {
    GameRoomEntity e = (GameRoomEntity) session.getAttributes().get(ROOM);
    if (e != null) {
      GameRoomSocketHandler room = rooms.get(e.getId());
      if (room != null) {
        room.handleMessage(message.getPayload());
      }
    }
  }
}
