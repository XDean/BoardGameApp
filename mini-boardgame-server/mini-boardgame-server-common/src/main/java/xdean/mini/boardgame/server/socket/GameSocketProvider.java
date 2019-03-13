package xdean.mini.boardgame.server.socket;

import org.springframework.web.socket.WebSocketSession;

import com.fasterxml.jackson.databind.JsonNode;

import io.reactivex.Observable;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

public interface GameSocketProvider {
  Observable<WebSocketEvent<?>> handle(WebSocketSession session, GameRoomEntity room, Observable<WebSocketEvent<JsonNode>> input);
}
