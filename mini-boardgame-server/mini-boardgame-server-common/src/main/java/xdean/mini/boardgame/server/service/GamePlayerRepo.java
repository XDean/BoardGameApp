package xdean.mini.boardgame.server.service;

import org.springframework.data.jpa.repository.JpaRepository;

import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;

public interface GamePlayerRepo extends JpaRepository<GamePlayerEntity, Integer> {

}
