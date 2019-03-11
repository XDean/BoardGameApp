package xdean.mini.boardgame.server.mybatis.mapper;

import java.util.Collection;
import java.util.Optional;

import org.apache.ibatis.annotations.Mapper;
import org.springframework.security.core.authority.SimpleGrantedAuthority;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;

@Mapper
public interface UserMapper {

  Optional<UserProfileEntity> findByUserUsername(String username);

  UserProfileEntity save(UserProfileEntity build);

  Optional<UserEntity> findByUsername(String username);

  Collection<SimpleGrantedAuthority> findAllByUsername(String username);

}
