package xdean.mini.boardgame.server.socket;

import lombok.Builder;
import lombok.Getter;

public class WebSocketIllegalArgumentException extends IllegalArgumentException {

  @Getter
  private final WebSocketEvent<?> event;

  @Builder
  public WebSocketIllegalArgumentException(String message, WebSocketEvent<?> event, Throwable cause) {
    super(message, cause);
    this.event = event;
  }
}
