package xdean.mini.boardgame.server.annotation.processor;

import java.util.List;

public class SocketEventDescription {

  String topic;

  public static class SocketEventSide{
    boolean fromServer;
    String desc;
    List<SocketEventAttr> attrs;
  }

  public static class SocketEventAttr{

  }
}
