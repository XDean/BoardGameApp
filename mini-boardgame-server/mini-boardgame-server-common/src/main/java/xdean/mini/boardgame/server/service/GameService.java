package xdean.mini.boardgame.server.service;

public interface GameService {
  String name();

  void createGame(int roomId);
}
