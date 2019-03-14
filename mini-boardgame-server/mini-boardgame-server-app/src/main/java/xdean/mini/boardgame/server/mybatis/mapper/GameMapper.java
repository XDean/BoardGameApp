package xdean.mini.boardgame.server.mybatis.mapper;

import java.util.List;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.session.RowBounds;

import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.mybatis.Tables;
import xdean.mybatis.extension.annotation.DefaultBuilder;

@Mapper
@DefaultBuilder(GameMapperBuilder.class)
public interface GameMapper extends BaseMapper {
  GamePlayerEntity findPlayer(int id);

  GameRoomEntity findRoom(int roomId);

  GameRoomEntity findRoomByPlayer(int playerId);

  List<GamePlayerEntity> findAllPlayersInRoom(int roomId);

  void savePlayer(GamePlayerEntity player);

  void saveRoom(GameRoomEntity room);

  void delete(GameRoomEntity room);

  List<GameRoomEntity> findAllRoom(String gameName, RowBounds page);

  default boolean roomExist(int id) {
    return exist(Tables.GameRoomTable.id, Integer.toString(id));
  }
}
