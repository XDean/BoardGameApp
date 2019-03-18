package xdean.mini.boardgame.server.annotation.processor.model;

import lombok.Builder;
import lombok.Value;

@Value
@Builder
public class SocketDescription {
  String topic;
  SocketSide fromServer;
  SocketSide fromClient;
}
