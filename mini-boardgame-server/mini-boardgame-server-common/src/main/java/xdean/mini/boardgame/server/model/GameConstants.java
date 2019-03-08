package xdean.mini.boardgame.server.model;

import xdean.mini.boardgame.server.annotation.Attr;
import xdean.mini.boardgame.server.annotation.BeanType;
import xdean.mini.boardgame.server.annotation.FromClient;
import xdean.mini.boardgame.server.annotation.FromServer;
import xdean.mini.boardgame.server.annotation.Payload;

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

    @BeanType(int.class)
    String FROM_SEAT = "FROM_SEAT";

    @BeanType(int.class)
    String TO_SEAT = "TO_SEAT";
  }

  interface SocketTopic {

    @FromServer(payload = @Payload(String.class))
    String ERROR_TOPIC = "ERROR";// an unexpected error happened

    @FromServer
    String BAD_CREDENTIAL = "BAD_CREDENTIAL";

    @FromClient(attr = @Attr(AttrKey.ACCESS_TOKEN))
    String AUTHENTICATION = "AUTHENTICATION";

    @FromServer(attr = @Attr(AttrKey.USER_ID))
    String PLAYER_JOIN = "PLAYER_JOIN";

    @FromServer(attr = @Attr(AttrKey.USER_ID))
    String PLAYER_EXIT = "PLAYER_EXIT";

    @FromServer(attr = @Attr(AttrKey.USER_ID))
    String PLAYER_CONNECT = "PLAYER_CONNECT";

    @FromServer(attr = @Attr(AttrKey.USER_ID))
    String PLAYER_DISCONNECT = "PLAYER_DISCONNECT";

    @FromServer
    String ROOM_CANCEL = "ROOM_CANCEL";

    @FromClient(attr = @Attr(value = AttrKey.TO_SEAT))
    @FromServer(attr = @Attr(value = AttrKey.FROM_SEAT))
    String CHANGE_SEAT_REQUEST = "CHANGE_SEAT_REQUEST";

    @FromServer(attr = @Attr(value = { AttrKey.FROM_SEAT, AttrKey.TO_SEAT }))
    String CHANGE_SEAT = "CHANGE_SEAT";

    @FromServer
    String GAME_START = "GAME_START";

    @FromServer(payload = @Payload(desc = "The game's information"))
    String GAME_INFO = "GAME_INFO";
  }
}
