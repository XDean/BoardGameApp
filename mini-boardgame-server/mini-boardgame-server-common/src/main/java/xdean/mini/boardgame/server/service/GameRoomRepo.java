package xdean.mini.boardgame.server.service;

import java.util.List;

import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

public interface GameRoomRepo extends JpaRepository<GameRoomEntity, Integer> {
  List<GameRoomEntity> findAllByRoomGameName(String gameName, Pageable pageable);
}
