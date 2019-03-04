package xdean.mini.boardgame.server.service;

import org.springframework.data.jpa.repository.JpaRepository;

import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

public interface GameRoomRepo extends JpaRepository<GameRoomEntity, Integer> {

}
