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

    @FromServer(
        desc = "An error happened",
        payload = @Payload(type = String.class, desc = "Error message"))
    String ERROR_TOPIC = "ERROR";

    @FromClient(
        desc = "Authenticate current socket connection.",
        attr = @Attr(desc = "Access token, can be find in cookie", value = AttrKey.ACCESS_TOKEN))
    String AUTHENTICATION = "AUTHENTICATION";

    @FromServer(desc = "The AUTHENTICATION is not valid")
    String BAD_CREDENTIAL = "BAD_CREDENTIAL";

    @FromServer(
        desc = "A player connect(after authentication) to this room",
        attr = @Attr(desc = "The connected player id", value = AttrKey.USER_ID))
    String PLAYER_CONNECT = "PLAYER_CONNECT";

    @FromServer(
        desc = "A player disconnect from this room",
        attr = @Attr(desc = "The disconnected player id", value = AttrKey.USER_ID))
    String PLAYER_DISCONNECT = "PLAYER_DISCONNECT";

    @FromServer(
        desc = "A player join into this room",
        attr = @Attr(desc = "The joined player id", value = AttrKey.USER_ID))
    String PLAYER_JOIN = "PLAYER_JOIN";

    @FromServer(
        desc = "A player exit this room",
        attr = @Attr(desc = "The exited player id", value = AttrKey.USER_ID))
    String PLAYER_EXIT = "PLAYER_EXIT";

    @FromServer(desc = "This room is canceled (all players exited)")
    String ROOM_CANCEL = "ROOM_CANCEL";

    @FromClient(
        desc = "To request change seat to specific seat. Server will forward the request to target player",
        attr = @Attr(desc = "The target seat number", value = AttrKey.TO_SEAT))
    @FromServer(
        desc = "A player is requesting to change seat with you",
        attr = @Attr(desc = "The seat of player who requst to change seat with you ", value = AttrKey.FROM_SEAT))
    String CHANGE_SEAT_REQUEST = "CHANGE_SEAT_REQUEST";

    @FromServer(
        desc = "A change seat request be accepted. There are 2 situations:\n"
            + "- Seat 1 player request to empty seat 2\n"
            + "- Seat 1 player request to seat 2 and seat 2 player accepted(send a request that change to 1)",
        attr = {
            @Attr(desc = "One of the changed seats", value = AttrKey.FROM_SEAT),
            @Attr(desc = "The other one of the changed seats", value = AttrKey.TO_SEAT) })
    String CHANGE_SEAT = "CHANGE_SEAT";

    @FromServer(desc = "Game Start!")
    String GAME_START = "GAME_START";

    @FromClient(desc = "To request this room's fully information")
    @FromServer(
        desc = "Response GAME_INFO request",
        payload = @Payload(desc = "The game's information. Different game has different info"))
    String GAME_INFO = "GAME_INFO";
  }
}
