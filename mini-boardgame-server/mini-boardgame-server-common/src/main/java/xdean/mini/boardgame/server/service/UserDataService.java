package xdean.mini.boardgame.server.service;

import java.util.Optional;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;

public interface UserDataService {
  Optional<UserEntity> getCurrentUser();

  Optional<UserEntity> findUserByUsername(String username);

  UserProfileEntity save(UserProfileEntity profile);
}
