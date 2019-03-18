package xdean.mini.boardgame.server.annotation.processor.model;

import javax.lang.model.type.TypeMirror;

import lombok.Builder;
import lombok.Value;

@Value
@Builder
public class SocketPayload {
  TypeMirror type;
  String desc;
}
