package xdean.mini.boardgame.server.model;

import xdean.mini.boardgame.server.annotation.Attrs;
import xdean.mini.boardgame.server.annotation.BeanType;

public interface GameConstants {

  interface TagKey {
  }

  interface ErrorCode {
    int ILLEGAL_ARGUEMENT = -100;
  }

  interface AttrKey {
    @BeanType(Integer.class)
    String USER_ID = "USER_ID";

    @BeanType(GameRoom.class)
    String ROOM = "ROOM";

    @BeanType(String.class)
    String ACCESS_TOKEN = "ACCESS_TOKEN";

    String FROM_SEAT = "FROM_SEAT";

    String TO_SEAT = "TO_SEAT";
  }

  interface SocketTopic {

    String ERROR_TOPIC = "ERROR";// an unexpected error happened

    String BAD_CREDENTIAL = "BAD_CREDENTIAL";

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

    String ROOM_CANCEL = "ROOM_CANCEL";

    @Attrs({ AttrKey.FROM_SEAT, AttrKey.TO_SEAT })
    String CHANGE_SEAT_REQUEST = "CHANGE_SEAT_REQUEST";

    @Attrs({ AttrKey.FROM_SEAT, AttrKey.TO_SEAT })
    String CHANGE_SEAT = "CHANGE_SEAT";
  }
}
