package xdean.mini.boardgame.server.model;

import xdean.mini.boardgame.server.annotation.Attrs;
import xdean.mini.boardgame.server.annotation.BeanType;

public interface GameConstants {

  interface TagKey {
  }

  interface AttrKey {
    @BeanType(Integer.class)
    String USER_ID = "USER_ID";

    @BeanType(GameRoom.class)
    String ROOM = "ROOM";

    @BeanType(String.class)
    String ACCESS_TOKEN = "ACCESS_TOKEN";
  }

  interface SocketTopic {

    @Attrs(AttrKey.ACCESS_TOKEN)
    String AUTHENTICATION = "AUTHENTICATION";

    @Attrs(AttrKey.USER_ID)
    String PLAYER_JOIN = "PLAYER_JOIN";

    @Attrs(AttrKey.USER_ID)
    String PLAYER_EXIT = "PLAYER_EXIT";

    @Attrs(AttrKey.USER_ID)
    String PLAYER_CONNECT = "PLAYER_CONNECT";

    @Attrs(AttrKey.USER_ID)
    String PLAYER_DISCONNECT = "PLAYER_DISCONNECT";
  }
}
