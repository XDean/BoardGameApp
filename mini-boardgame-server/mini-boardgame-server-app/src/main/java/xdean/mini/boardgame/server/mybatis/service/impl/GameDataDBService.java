package xdean.mini.boardgame.server.mybatis.service.impl;

import java.util.List;
import java.util.Optional;

import javax.inject.Inject;

import org.springframework.stereotype.Service;

import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.model.param.SearchGameRequest;
import xdean.mini.boardgame.server.mybatis.mapper.GameMapper;
import xdean.mini.boardgame.server.service.GameDataService;

@Service
public class GameDataDBService implements GameDataService {

  @Inject
  GameMapper gameMapper;

  @Override
  public GamePlayerEntity findPlayer(int id) {
    GameRoomEntity room = gameMapper.findRoomByPlayer(id);
    if (room != null) {
      List<GamePlayerEntity> players = gameMapper.findAllPlayersInRoom(room.getId());
      room.setPlayers(players);
      players.forEach(e -> e.setRoom(room));
      GamePlayerEntity player = players.stream()
          .filter(e -> e.getId() == id)
          .findAny()
          .orElse(null);
      if (player != null) {
        return player;
      }
    }
    return GamePlayerEntity.builder().id(id).build();
  }

  @Override
  public Optional<GameRoomEntity> findRoom(int roomId) {
    GameRoomEntity room = gameMapper.findRoom(roomId);
    if (room != null) {
      fillRoom(room);
      return Optional.of(room);
    }
    return Optional.empty();
  }

  @Override
  public void save(GamePlayerEntity player) {
    gameMapper.savePlayer(player);
    player.getRoom().ifPresent(this::save);
  }

  @Override
  public void save(GameRoomEntity room) {
    gameMapper.saveRoom(room);
    room.getPlayers().forEach(e -> gameMapper.savePlayer(e));
  }

  @Override
  public void delete(GameRoomEntity room) {
    gameMapper.delete(room.getId());
  }

  @Override
  public List<GameRoomEntity> searchGame(SearchGameRequest request) {
    List<GameRoomEntity> rooms = gameMapper.searchRoom(request);
    rooms.forEach(this::fillRoom);
    return rooms;
  }

  @Override
  public void saveAll(List<GamePlayerEntity> players) {
    players.forEach(e -> gameMapper.savePlayer(e));
  }

  @Override
  public boolean roomExist(int id) {
    return gameMapper.roomExist(id);
  }

  private void fillRoom(GameRoomEntity room) {
    List<GamePlayerEntity> players = gameMapper.findAllPlayersInRoom(room.getId());
    room.setPlayers(players);
    room.getBoard().setRoom(room);
    players.forEach(e -> e.setRoom(room));
  }
}
