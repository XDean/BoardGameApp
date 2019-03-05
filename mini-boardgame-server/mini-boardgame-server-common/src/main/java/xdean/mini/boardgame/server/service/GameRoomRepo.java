package xdean.mini.boardgame.server.service;

import java.util.List;
import java.util.Optional;

import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

import xdean.mini.boardgame.server.model.entity.GameRoomEntity;

public interface GameRoomRepo extends JpaRepository<GameRoomEntity, Integer> {
  List<GameRoomEntity> findAllByRoomGameName(String gameName, Pageable pageable);

  Optional<GameRoomEntity> findByPlayersUserId(int userId);
}
