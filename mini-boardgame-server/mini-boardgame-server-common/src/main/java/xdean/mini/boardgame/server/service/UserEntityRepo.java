package xdean.mini.boardgame.server.service;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

import xdean.mini.boardgame.server.model.entity.UserEntity;

public interface UserEntityRepo extends JpaRepository<UserEntity, Integer> {
  Optional<UserEntity> findByUsername(String username);
}
