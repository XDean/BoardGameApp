package xdean.mini.boardgame.server.mybatis.mapper;

import java.util.List;

import org.apache.ibatis.annotations.Mapper;

import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.model.param.SearchGameRequest;
import xdean.mini.boardgame.server.mybatis.Tables;
import xdean.mybatis.extension.annotation.DefaultBuilder;
import xdean.mybatis.extension.annotation.ResultMapType;

@Mapper
@DefaultBuilder(GameMapperBuilder.class)
public interface GameMapper extends BaseMapper {
  @ResultMapType
  GamePlayerEntity findPlayer(Integer id);

  @ResultMapType
  GameRoomEntity findRoom(Integer roomId);

  @ResultMapType
  GameRoomEntity findRoomByPlayer(Integer playerId);

  @ResultMapType
  List<GamePlayerEntity> findAllPlayersInRoom(Integer roomId);

  void savePlayer(GamePlayerEntity player);

  void saveRoom(GameRoomEntity room);

  void delete(Integer roomId);

  @ResultMapType
  List<GameRoomEntity> searchRoom(SearchGameRequest request);

  default boolean roomExist(int id) {
    return exist(Tables.GameRoomTable.id, Integer.toString(id));
  }
}
