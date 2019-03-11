package xdean.mini.boardgame.server.mybatis.mapper;

import java.util.List;
import java.util.Optional;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.session.RowBounds;

import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

@Mapper
public interface GameMapper {

  Optional<GamePlayerEntity> findPlayer(int id);

  Optional<GameRoomEntity> findRoom(int roomId);

  void save(GamePlayerEntity player);

  void save(GameRoomEntity room);

  void delete(GameRoomEntity room);

  List<GameRoomEntity> findAllRoom(String gameName, RowBounds page);
}
