package xdean.mini.boardgame.server.mybatis.service.impl;

import java.util.Optional;

import javax.inject.Inject;

import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mini.boardgame.server.mybatis.mapper.UserMapper;
import xdean.mini.boardgame.server.service.UserDataService;

@Service
public class UserDataDBService implements UserDataService {

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
  @Transactional
  public Optional<UserEntity> findUserByUsername(String username) {
    UserEntity user = userMapper.findByUsername(username);
    if (user != null) {
      user.setAuthorities(userMapper.findAuthorities(user.getId()));
    }
    return Optional.ofNullable(user);
  }

  @Override
  public void save(UserProfileEntity profile) {
    userMapper.save(profile);
  }

  @Override
  public void save(UserEntity user) {
    userMapper.save(user);
  }

  @Override
  public void delete(String username) {
    userMapper.delete(username);
  }

  @Override
  public void changePassword(String username, String password) {
    userMapper.changePassword(username, password);
  }

  @Override
  public boolean userExist(int id) {
    return userMapper.userExist(id);
  }

  @Override
  public boolean userExist(String username) {
    return userMapper.userExist(username);
  }
}
