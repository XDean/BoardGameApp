package xdean.mini.boardgame.server.service.impl;

import java.util.Collections;
import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.Random;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import javax.inject.Inject;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.stereotype.Service;
import org.springframework.util.Assert;
import org.springframework.web.socket.WebSocketSession;

import com.fasterxml.jackson.databind.JsonNode;

import io.reactivex.Observable;
import io.reactivex.subjects.PublishSubject;
import io.reactivex.subjects.Subject;
import xdean.mini.boardgame.server.model.GameConstants;
import xdean.mini.boardgame.server.model.GameConstants.SocketTopic;
import xdean.mini.boardgame.server.model.GameRoom;
import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.param.CreateGameRequest;
import xdean.mini.boardgame.server.model.param.CreateGameResponse;
import xdean.mini.boardgame.server.model.param.CurrentGameResponse;
import xdean.mini.boardgame.server.model.param.ExitGameRequest;
import xdean.mini.boardgame.server.model.param.ExitGameResponse;
import xdean.mini.boardgame.server.model.param.GameCenterErrorCode;
import xdean.mini.boardgame.server.model.param.JoinGameRequest;
import xdean.mini.boardgame.server.model.param.JoinGameResponse;
import xdean.mini.boardgame.server.model.param.SearchGameRequest;
import xdean.mini.boardgame.server.model.param.SearchGameResponse;
import xdean.mini.boardgame.server.service.GameCenterService;
import xdean.mini.boardgame.server.service.GamePlayerRepo;
import xdean.mini.boardgame.server.service.GameRoomRepo;
import xdean.mini.boardgame.server.service.GameService;
import xdean.mini.boardgame.server.service.UserService;
import xdean.mini.boardgame.server.socket.WebSocketEvent;
import xdean.mini.boardgame.server.socket.WebSocketEvent.WebSocketEventBuilder;
import xdean.mini.boardgame.server.socket.WebSocketProvider;
import xdean.mini.boardgame.server.socket.WebSocketSendType;
import xdean.mini.boardgame.server.util.JpaUtil;

@Service
public class GameCenterServiceImpl implements GameCenterService, WebSocketProvider {

  @Autowired(required = false)
  List<GameService> games = Collections.emptyList();

  private @Inject UserService userService;
  private @Inject GamePlayerRepo gamePlayerRepo;
  private @Inject GameRoomRepo gameRoomRepo;

  private final Object[] locks = IntStream.range(0, 32).mapToObj(i -> new Object()).toArray();

  @Override
  public CreateGameResponse createGame(CreateGameRequest request) {
    Optional<GameService> game = findGame(request.getGameName());
    if (!game.isPresent()) {
      return CreateGameResponse.builder()
          .errorCode(GameCenterErrorCode.NO_SUCH_GAME)
          .build();
    }
    Optional<UserEntity> user = userService.getCurrentUser();
    if (!user.isPresent()) {
      return CreateGameResponse.builder()
          .errorCode(GameCenterErrorCode.NO_USER)
          .build();
    }
    UserEntity e = user.get();
    synchronized (getLock(e.getId())) {
      GamePlayerEntity player = JpaUtil.findOrCreate(gamePlayerRepo, e.getId(),
          id -> GamePlayerEntity.builder().userId(id).build());
      if (player.getRoom() != null) {
        return CreateGameResponse.builder()
            .errorCode(GameCenterErrorCode.ALREADY_IN_ROOM)
            .build();
      }
      Integer roomId = generateId();
      GameRoomEntity room = GameRoomEntity.builder()
          .id(roomId)
          .room(GameRoom.builder()
              .gameName(request.getGameName())
              .createdTime(new Date())
              .playerCount(request.getPlayerCount())
              .roomName(request.getRoomName().isEmpty() ? "Room " + roomId : request.getRoomName())
              .build())
          .player(player)
          .build();
      player.setRoom(room);
      room = gameRoomRepo.save(room);
      // gamePlayerRepo.save(player);
      return CreateGameResponse.builder()
          .roomId(roomId)
          .build();
    }
  }

  @Override
  public JoinGameResponse joinGame(JoinGameRequest request) {
    Optional<UserEntity> user = userService.getCurrentUser();
    if (!user.isPresent()) {
      return JoinGameResponse.builder()
          .errorCode(GameCenterErrorCode.NO_USER)
          .build();
    }
    synchronized (getLock(user.get().getId())) {
      synchronized (getLock(request.getRoomId())) {
        Optional<GameRoomEntity> oRoom = gameRoomRepo.findById(request.getRoomId());
        if (!oRoom.isPresent()) {
          return JoinGameResponse.builder()
              .errorCode(GameCenterErrorCode.NO_SUCH_ROOM)
              .build();
        }
        GameRoomEntity room = oRoom.get();
        GamePlayerEntity player = JpaUtil.findOrCreate(gamePlayerRepo, user.get().getId(),
            id -> GamePlayerEntity.builder().userId(id).build());
        if (player.getRoom() != null) {
          return JoinGameResponse.builder()
              .errorCode(GameCenterErrorCode.ALREADY_IN_ROOM)
              .build();
        }
        room.addPlayer(player);
        player.setRoom(room);
        gameRoomRepo.save(room);
        gamePlayerRepo.save(player);
        sendEvent(room.getId(), player.getUserId(), SocketTopic.PLAYER_JOIN);
        return JoinGameResponse.builder()
            .build();
      }
    }
  }

  @Override
  public ExitGameResponse exitGame(ExitGameRequest request) {
    Optional<UserEntity> user = userService.getCurrentUser();
    if (!user.isPresent()) {
      return ExitGameResponse.builder()
          .errorCode(GameCenterErrorCode.NO_USER)
          .build();
    }
    synchronized (getLock(user.get().getId())) {
      GamePlayerEntity player = JpaUtil.findOrCreate(gamePlayerRepo, user.get().getId(),
          id -> GamePlayerEntity.builder().userId(id).build());
      GameRoomEntity room = player.getRoom();
      if (room == null) {
        return ExitGameResponse.builder()
            .errorCode(GameCenterErrorCode.NOT_IN_ROOM)
            .build();
      }
      synchronized (getLock(room.getId())) {
        player.setRoom(null);
        gamePlayerRepo.save(player);
        room.removePlayer(player);
        if (room.getPlayers().isEmpty()) {
          gameRoomRepo.delete(room);
        } else {
          gameRoomRepo.save(room);
        }
        sendEvent(room.getId(), player.getUserId(), SocketTopic.PLAYER_EXIT);
        if (room.getPlayers().isEmpty()) {
          sendEvent(room.getId(), null, SocketTopic.ROOM_CANCEL);
        }
        return ExitGameResponse.builder().build();
      }
    }
  }

  @Override
  public SearchGameResponse searchGame(SearchGameRequest request) {
    Optional<GameService> game = findGame(request.getGameName());
    if (!game.isPresent()) {
      return SearchGameResponse.builder()
          .errorCode(GameCenterErrorCode.NO_SUCH_GAME)
          .build();
    }
    List<GameRoomEntity> rooms = gameRoomRepo.findAllByRoomGameName(request.getGameName(),
        PageRequest.of(request.getPage(), request.getPageSize()));
    return SearchGameResponse.builder()
        .rooms(rooms.stream()
            // .peek(e ->
            // e.getRoom().setCurrentPlayerCount(e.getPlayers().size()))
            .map(e -> e.getRoom())
            .collect(Collectors.toList()))
        .build();
  }

  @Override
  public CurrentGameResponse currentGame() {
    Optional<UserEntity> user = userService.getCurrentUser();
    if (!user.isPresent()) {
      return CurrentGameResponse.builder()
          .errorCode(GameCenterErrorCode.NO_USER)
          .build();
    }
    synchronized (getLock(user.get().getId())) {
      GamePlayerEntity player = JpaUtil.findOrCreate(gamePlayerRepo, user.get().getId(),
          id -> GamePlayerEntity.builder().userId(id).build());
      GameRoomEntity room = player.getRoom();
      if (room == null) {
        return CurrentGameResponse.builder()
            .errorCode(GameCenterErrorCode.NOT_IN_ROOM)
            .build();
      }
      return CurrentGameResponse.builder().room(room.getRoom()).build();
    }
  }

  Map<Integer, Subject<WebSocketEvent<?>>> roomSubjects = new HashMap<>();

  @Override
  public Observable<WebSocketEvent<?>> handle(WebSocketSession session, GameRoom room,
      Observable<WebSocketEvent<JsonNode>> input) {
    Integer id = (Integer) session.getAttributes().get(GameConstants.AttrKey.USER_ID);
    Assert.notNull(id, "Authed user must have id");
    Subject<WebSocketEvent<?>> subject = roomSubjects.computeIfAbsent(room.getId(), r -> PublishSubject.create());
    subject.onNext(WebSocketEvent.builder()
        .type(WebSocketSendType.SELF)
        .topic(SocketTopic.PLAYER_CONNECT)
        .attribute(GameConstants.AttrKey.USER_ID, id)
        .build());
    input.subscribe(e -> {
    }, e -> subject.onError(e), () -> subject.onNext(WebSocketEvent.builder()
        .type(WebSocketSendType.SELF)
        .topic(SocketTopic.PLAYER_DISCONNECT)
        .attribute(GameConstants.AttrKey.USER_ID, id)
        .build()));
    return subject;
  }

  private void sendEvent(int roomId, Integer playerId, String topic) {
    Subject<WebSocketEvent<?>> subject = roomSubjects.get(roomId);
    if (subject != null) {
      WebSocketEventBuilder<?> builder = WebSocketEvent.builder()
          .type(WebSocketSendType.SELF)
          .topic(topic);
      if (playerId != null) {
        builder.attribute(GameConstants.AttrKey.USER_ID, playerId);
      }
      subject.onNext(builder.build());
    }
  }

  private Integer generateId() {
    Random r = new Random();
    Integer id;
    do {
      id = r.nextInt(1000000);
    } while (gameRoomRepo.existsById(id));
    return id;
  }

  private Optional<GameService> findGame(String name) {
    return games.stream().filter(g -> g.name().equals(name)).findFirst();
  }

  private Object getLock(int id) {
    return locks[id % 32];
  }
}
