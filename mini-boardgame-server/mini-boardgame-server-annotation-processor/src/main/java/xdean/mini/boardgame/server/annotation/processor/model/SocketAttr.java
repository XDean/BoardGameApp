package xdean.mini.boardgame.server.annotation.processor.model;

import lombok.Builder;
import lombok.Value;

@Value
@Builder(toBuilder = true)
public class SocketAttr {
  String key;
  String type;
  String desc;
}