package xdean.mini.boardgame.server.service;

import java.util.List;
import java.util.Optional;

import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.model.param.SearchGameRequest;

public interface GameDataService {

  GamePlayerEntity findPlayer(int id);

  Optional<GameRoomEntity> findRoom(int roomId);

  void save(GamePlayerEntity player);

  void save(GameRoomEntity room);

  void delete(GameRoomEntity room);

  List<GameRoomEntity> searchGame(SearchGameRequest request);

  void saveAll(List<GamePlayerEntity> asList);

  boolean roomExist(int id);
}
