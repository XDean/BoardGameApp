package xdean.mini.boardgame.server.annotation.processor.model;

import javax.lang.model.type.TypeMirror;

import lombok.Builder;
import lombok.Value;

@Value
@Builder(toBuilder = true)
public class SocketAttr {
  String key;
  TypeMirror type;
  String desc;
}