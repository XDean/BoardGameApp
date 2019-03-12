package xdean.mini.boardgame.server.service.impl;

import java.util.Optional;

import javax.inject.Inject;

import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.stereotype.Service;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mini.boardgame.server.mybatis.mapper.UserMapper;
import xdean.mini.boardgame.server.service.UserDataService;

@Service
public class UserServiceImpl implements UserDataService {

  @Inject
  UserMapper userMapper;

  @Override
  public Optional<UserEntity> getCurrentUser() {
    Authentication a = SecurityContextHolder.getContext().getAuthentication();
    if (a != null) {
      Object p = a.getPrincipal();
      if (p instanceof User) {
        String username = ((User) p).getUsername();
        return findUserByUsername(username);
      }
    }
    return Optional.empty();
  }

  @Override
  public Optional<UserEntity> findUserByUsername(String username) {
    return Optional.ofNullable(userMapper.findByUsername(username));
  }

  @Override
  public UserProfileEntity save(UserProfileEntity profile) {
    userMapper.save(profile);
    return profile;
  }
}
