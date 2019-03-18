package xdean.mini.boardgame.server.annotation.processor.model;

import java.util.List;

import javax.annotation.CheckForNull;

import lombok.Builder;
import lombok.Singular;
import lombok.Value;

@Value
@Builder
public class SocketSide {
  boolean fromServer;
  String desc;

  @Singular
  List<SocketAttr> attrs;

  @CheckForNull
  SocketPayload payload;
}