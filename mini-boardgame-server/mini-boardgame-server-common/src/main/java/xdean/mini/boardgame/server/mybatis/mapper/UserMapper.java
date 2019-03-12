package xdean.mini.boardgame.server.mybatis.mapper;

import org.apache.ibatis.annotations.Mapper;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mybatis.extension.annotation.DefaultBuilder;

@Mapper
@DefaultBuilder(UserMapperBuilder.class)
public interface UserMapper {
  void save(UserProfileEntity profile);

  UserEntity findByUsername(String username);
}
