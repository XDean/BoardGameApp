package xdean.mini.boardgame.server.service;

import java.util.Optional;

import xdean.mini.boardgame.server.model.entity.UserEntity;

public interface UserService {
  Optional<UserEntity> getCurrentUser();

  Optional<UserEntity> getUserByUsername(String username);
}
