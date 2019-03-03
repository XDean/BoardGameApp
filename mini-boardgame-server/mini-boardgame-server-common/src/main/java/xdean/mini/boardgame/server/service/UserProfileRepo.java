package xdean.mini.boardgame.server.service;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

import xdean.mini.boardgame.server.model.entity.UserProfileEntity;

public interface UserProfileRepo extends JpaRepository<UserProfileEntity, Integer> {

  Optional<UserProfileEntity> findByUserUsername(String username);
}
