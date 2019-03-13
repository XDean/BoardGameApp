package xdean.mini.boardgame.server.mybatis.mapper;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mini.boardgame.server.mybatis.Tables;
import xdean.mybatis.extension.resultmap.InitResultMap;

@Configuration
public class UserResultMap implements Tables {

  @Bean
  public InitResultMap<UserProfileEntity> userProfileMapper() {
    return InitResultMap.create(UserProfileEntity.class)
        .namespace()
        .id(UserProfileEntity.class.getName())
        .resultMap(b -> b
            .stringFree()
            .mapping(ProfileTable.id, UserProfileEntity::setId)
            .mapping(ProfileTable.nickname, UserProfileEntity::setNickname)
            .mapping(ProfileTable.male, UserProfileEntity::setMale)
            .mapping(ProfileTable.avatarUrl, UserProfileEntity::setAvatarUrl))
        .build();
  }

  @Bean
  public InitResultMap<UserEntity> userEntityMapper() {
    return InitResultMap.create(UserEntity.class)
        .namespace()
        .id(UserEntity.class.getName())
        .resultMap(b -> b
            .stringFree()
            .mapping(UserTable.id, UserEntity::setId)
            .mapping(UserTable.username, UserEntity::setUsername)
            .mapping(UserTable.password, UserEntity::setPassword)
            .mapping(UserTable.enabled, UserEntity::setEnabled)
            .mapping(d -> d.property(UserEntity::setProfile).nestMap(userProfileMapper().getId())))
        .build();
  }
}
