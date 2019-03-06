package xdean.mini.boardgame.server.model;

import xdean.mini.boardgame.server.annotation.Attrs;
import xdean.mini.boardgame.server.annotation.BeanType;
import xdean.mini.boardgame.server.annotation.Tag;

public interface GameConstants {

  interface TagKey {
    String SOCKET_ATTR = "SOCKET_ATTR";
  }

  interface AttrKey {
    @Tag(TagKey.SOCKET_ATTR)
    @BeanType(Integer.class)
    String PLAYER_ID = "PLAYER_ID";

    @Tag(TagKey.SOCKET_ATTR)
    @BeanType(GameRoom.class)
    String ROOM = "ROOM";
  }

  interface SocketTopic {
    @Attrs(AttrKey.PLAYER_ID)
    String PLAYER_JOIN = "PLAYER_JOIN";

    @Attrs(AttrKey.PLAYER_ID)
    String PLAYER_EXIT = "PLAYER_EXIT";

    @Attrs(AttrKey.PLAYER_ID)
    String PLAYER_CONNECT = "PLAYER_CONNECT";

    @Attrs(AttrKey.PLAYER_ID)
    String PLAYER_DISCONNECT = "PLAYER_DISCONNECT";
  }
}
