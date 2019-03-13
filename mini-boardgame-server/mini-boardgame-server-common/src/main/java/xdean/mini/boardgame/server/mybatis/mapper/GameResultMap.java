package xdean.mini.boardgame.server.mybatis.mapper;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import xdean.mini.boardgame.server.model.Tables;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mybatis.extension.ConfigurationInitializer;
import xdean.mybatis.extension.resultmap.InitResultMap;

@Configuration
public class GameResultMap implements Tables {

  @Bean
  public ConfigurationInitializer gameRoomMapper() {
    return InitResultMap.create(GameRoomEntity.class)
        .namespace()
        .id(GameRoomEntity.class.getName())
        .resultMap(b -> b
            .stringFree()
            .mapping(GameRoomTable.id, GameRoomEntity::setId)
            .mapping(GameRoomTable.gameName, GameRoomEntity::setGameName)
            .mapping(GameRoomTable.roomName, GameRoomEntity::setRoomName)
            .mapping(GameRoomTable.createdTime, GameRoomEntity::setCreatedTime)
            .mapping(GameRoomTable.playerCount, GameRoomEntity::setPlayerCount))
        .build();
  }
}
