package xdean.mini.boardgame.server.mybatis.mapper;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.model.handler.GameBoardConverter;
import xdean.mini.boardgame.server.mybatis.Tables;
import xdean.mybatis.extension.resultmap.InitResultMap;

@Configuration
public class GameResultMap implements Tables {

  @Bean
  public InitResultMap<GameRoomEntity> gameRoomMapper() {
    return InitResultMap.create(GameRoomEntity.class)
        .namespace()
        .id(GameRoomEntity.class.getName())
        .resultMap(b -> b
            .stringFree()
            .mapping(GameRoomTable.id, GameRoomEntity::setId)
            .mapping(GameRoomTable.gameName, GameRoomEntity::setGameName)
            .mapping(GameRoomTable.roomName, GameRoomEntity::setRoomName)
            .mapping(GameRoomTable.createdTime, GameRoomEntity::setCreatedTime)
            .mapping(GameRoomTable.playerCount, GameRoomEntity::setPlayerCount)
            .mapping(d -> d.column(GameRoomTable.board).property(GameRoomEntity::setBoard).typeHandler(GameBoardConverter.class)))
        .build();
  }

  @Bean
  public InitResultMap<GamePlayerEntity> gamePlayerMapper() {
    return InitResultMap.create(GamePlayerEntity.class)
        .namespace()
        .id(GamePlayerEntity.class.getName())
        .resultMap(b -> b
            .stringFree()
            .mapping(GamePlayerTable.id, GamePlayerEntity::setUserId)
            .mapping(GamePlayerTable.seat, GamePlayerEntity::setSeat))
        .build();
  }
}
