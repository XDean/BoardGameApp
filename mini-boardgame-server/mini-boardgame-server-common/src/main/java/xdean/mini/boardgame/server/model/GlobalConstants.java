package xdean.mini.boardgame.server.model;

import xdean.mini.boardgame.server.annotation.Attr;
import xdean.mini.boardgame.server.annotation.Payload;
import xdean.mini.boardgame.server.annotation.Side;
import xdean.mini.boardgame.server.annotation.Topic;
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
    interface Category {
      String DEFAULT = "Default";
      String GAME_ROOM = "Game Room";
    }

    @Topic(category = Category.DEFAULT,
        fromServer = @Side(desc = "An error happened", payload = @Payload(type = String.class, desc = "Error message")))
    String ERROR_TOPIC = "ERROR";

    @Topic(category = Category.DEFAULT,
        fromClient = @Side(
            desc = "Authenticate current socket connection.",
            attr = @Attr(desc = "Access token, can be find in cookie", value = AttrKey.ACCESS_TOKEN)))
    String AUTHENTICATION = "AUTHENTICATION";

    @Topic(category = Category.DEFAULT,
        fromServer = @Side(desc = "The AUTHENTICATION is not valid"))
    String BAD_CREDENTIAL = "BAD_CREDENTIAL";

    @Topic(category = Category.GAME_ROOM,
        fromServer = @Side(
            desc = "A player connect(after authentication) to this room",
            attr = @Attr(desc = "The connected player id", value = AttrKey.USER_ID)))
    String PLAYER_CONNECT = "PLAYER_CONNECT";

    @Topic(category = Category.GAME_ROOM,
        fromServer = @Side(
            desc = "A player disconnect from this room",
            attr = @Attr(desc = "The disconnected player id", value = AttrKey.USER_ID)))
    String PLAYER_DISCONNECT = "PLAYER_DISCONNECT";

    @Topic(category = Category.GAME_ROOM,
        fromServer = @Side(
            desc = "A player join into this room",
            attr = @Attr(desc = "The joined player id", value = AttrKey.USER_ID)))
    String PLAYER_JOIN = "PLAYER_JOIN";

    @Topic(category = Category.GAME_ROOM,
        fromServer = @Side(
            desc = "A player exit this room",
            attr = @Attr(desc = "The exited player id", value = AttrKey.USER_ID)))
    String PLAYER_EXIT = "PLAYER_EXIT";

    @Topic(category = Category.GAME_ROOM,
        fromServer = @Side(desc = "This room is canceled (all players exited)"))
    String ROOM_CANCEL = "ROOM_CANCEL";

    @Topic(category = Category.GAME_ROOM,
        fromClient = @Side(
            desc = "To request change seat to specific seat. Server will forward the request to target player",
            attr = @Attr(desc = "The target seat number", value = AttrKey.TO_SEAT)),
        fromServer = @Side(
            desc = "A player is requesting to change seat with you",
            attr = @Attr(desc = "The seat of player who requst to change seat with you ", value = AttrKey.FROM_SEAT)))
    String CHANGE_SEAT_REQUEST = "CHANGE_SEAT_REQUEST";

    @Topic(category = Category.GAME_ROOM,
        fromServer = @Side(
            desc = "A change seat request be accepted. There are 2 situations:\n"
                + "- Seat 1 player request to empty seat 2\n"
                + "- Seat 1 player request to seat 2 and seat 2 player accepted(send a request that change to 1)",
            attr = {
                @Attr(desc = "One of the changed seats", value = AttrKey.FROM_SEAT),
                @Attr(desc = "The other one of the changed seats", value = AttrKey.TO_SEAT) }))
    String CHANGE_SEAT = "CHANGE_SEAT";

    @Topic(category = Category.GAME_ROOM,
        fromServer = @Side(desc = "Game Start!"))
    String GAME_START = "GAME_START";

    @Topic(category = Category.GAME_ROOM,
        fromClient = @Side(desc = "To request this room's fully information"),
        fromServer = @Side(
            desc = "Response GAME_INFO request",
            payload = @Payload(desc = "The game's information. Different game has different info")))
    String GAME_INFO = "GAME_INFO";
  }
}
