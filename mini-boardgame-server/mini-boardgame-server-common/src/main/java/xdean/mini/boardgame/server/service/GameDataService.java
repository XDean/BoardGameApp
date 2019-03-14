package xdean.mini.boardgame.server.service;

import java.util.List;
import java.util.Optional;

import org.apache.ibatis.session.RowBounds;

import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

public interface GameDataService {

  GamePlayerEntity findPlayer(int id);

  Optional<GameRoomEntity> findRoom(int roomId);

  void save(GamePlayerEntity player);

  void save(GameRoomEntity room);

  void delete(GameRoomEntity room);

  List<GameRoomEntity> findAllByRoomGameName(String gameName, RowBounds page);

  void saveAll(List<GamePlayerEntity> asList);

  boolean roomExist(int id);
}
