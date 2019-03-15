package xdean.mini.boardgame.server.mybatis.mapper;

import java.util.List;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.session.RowBounds;

import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.mybatis.Tables;
import xdean.mybatis.extension.annotation.DefaultBuilder;
import xdean.mybatis.extension.annotation.ResultMapType;

@Mapper
@DefaultBuilder(GameMapperBuilder.class)
public interface GameMapper extends BaseMapper {
  @ResultMapType(GamePlayerEntity.class)
  GamePlayerEntity findPlayer(Integer id);

  @ResultMapType(GameRoomEntity.class)
  GameRoomEntity findRoom(Integer roomId);

  @ResultMapType(GameRoomEntity.class)
  GameRoomEntity findRoomByPlayer(Integer playerId);

  @ResultMapType(GamePlayerEntity.class)
  List<GamePlayerEntity> findAllPlayersInRoom(Integer roomId);

  void savePlayer(GamePlayerEntity player);

  void saveRoom(GameRoomEntity room);

  void delete(GameRoomEntity room);

  @ResultMapType(GameRoomEntity.class)
  List<GameRoomEntity> findAllRoom(String gameName, RowBounds page);

  default boolean roomExist(int id) {
    return exist(Tables.GameRoomTable.id, Integer.toString(id));
  }
}
