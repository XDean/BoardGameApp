package xdean.mini.boardgame.server.socket;

import org.springframework.web.socket.WebSocketSession;

import com.fasterxml.jackson.databind.JsonNode;

import io.reactivex.Observable;
import xdean.mini.boardgame.server.model.GameRoom;

public interface WebSocketProvider {
  Observable<WebSocketEvent<?>> handle(WebSocketSession session, GameRoom room, Observable<WebSocketEvent<JsonNode>> input);
}
