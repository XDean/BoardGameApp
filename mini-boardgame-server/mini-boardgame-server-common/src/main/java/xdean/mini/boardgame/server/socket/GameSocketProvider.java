package xdean.mini.boardgame.server.socket;

import io.reactivex.Observable;

public interface GameSocketProvider {
  Observable<WebSocketEvent> handle(Observable<WebSocketEvent> input);
}
