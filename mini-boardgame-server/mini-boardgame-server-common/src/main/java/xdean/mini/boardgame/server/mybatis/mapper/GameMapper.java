package xdean.mini.boardgame.server.mybatis.mapper;

import java.util.List;
import java.util.Optional;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.session.RowBounds;

import io.reactivex.functions.IntFunction;
import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.model.entity.UserEntity;

@Mapper
public interface GameMapper {

  Optional<GameRoomEntity> findById(int roomId);

  Optional<GameRoomEntity> findByPlayersUserId(int id);

  GamePlayerEntity findOrCreateById(int id, IntFunction<GamePlayerEntity> object);

  GameRoomEntity save(GameRoomEntity room);

  void delete(GameRoomEntity room);

  void save(GamePlayerEntity player);

  List<GameRoomEntity> findAllByRoomGameName(String gameName, RowBounds page);

  void saveAll(List<GamePlayerEntity> asList);

  boolean existsById(Integer id);

}
