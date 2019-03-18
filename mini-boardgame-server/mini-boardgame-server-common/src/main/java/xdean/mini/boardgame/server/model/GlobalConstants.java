package xdean.mini.boardgame.server.model;

import xdean.mini.boardgame.server.annotation.Attr;
import xdean.mini.boardgame.server.annotation.FromClient;
import xdean.mini.boardgame.server.annotation.FromServer;
import xdean.mini.boardgame.server.annotation.Payload;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

public interface GlobalConstants {

  interface TagKey {
  }

  interface ErrorCode {
    int SUCCESS = 0;
    int ILLEGAL_ARGUEMENT = -100;
  }

  interface AttrKey {
    @Attr(type = Integer.class)
    String USER_ID = "USER_ID";

    @Attr(type = GameRoomEntity.class)
    String ROOM = "ROOM";

    @Attr(type = String.class)
    String ACCESS_TOKEN = "ACCESS_TOKEN";

    @Attr(type = int.class)
    String FROM_SEAT = "FROM_SEAT";

    @Attr(type = int.class)
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

    @FromClient(attr = @Attr(AttrKey.TO_SEAT))
    @FromServer(attr = @Attr(AttrKey.FROM_SEAT))
    String CHANGE_SEAT_REQUEST = "CHANGE_SEAT_REQUEST";

    @FromServer(attr = { @Attr(AttrKey.FROM_SEAT), @Attr(AttrKey.TO_SEAT) })
    String CHANGE_SEAT = "CHANGE_SEAT";

    @FromServer
    String GAME_START = "GAME_START";

    @FromServer(payload = @Payload(desc = "The game's information"))
    String GAME_INFO = "GAME_INFO";
  }
}
