package xdean.mini.boardgame.server.model.param;

public interface GameCenterErrorCode {
  static int NO_USER = 1;
  static int NO_SUCH_GAME = 2;
  static int NO_SUCH_ROOM = 3;
  static int ROOM_FULL = 4;
  static int ALREADY_IN_ROOM = 5;
  static int NOT_IN_ROOM = 6;
}
