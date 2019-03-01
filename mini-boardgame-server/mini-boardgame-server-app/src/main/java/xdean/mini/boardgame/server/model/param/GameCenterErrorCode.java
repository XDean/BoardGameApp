package xdean.mini.boardgame.server.model.param;

public interface GameCenterErrorCode {
  static int NO_SUCH_GAME = 1;
  static int NO_SUCH_ROOM = 2;
  static int ROOM_FULL = 3;
  static int ALREADY_IN_ROOM = 4;
  static int NOT_IN_ROOM = 5;
}
