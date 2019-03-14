package xdean.mini.boardgame.server.service;

import java.util.Optional;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;

public interface UserDataService {
  Optional<UserEntity> getCurrentUser();

  Optional<UserEntity> findUserByUsername(String username);

  void save(UserEntity user);

  void save(UserProfileEntity profile);

  void delete(String username);

  void changePassword(String username, String password);

  boolean userExist(int id);

  boolean userExist(String username);
}
