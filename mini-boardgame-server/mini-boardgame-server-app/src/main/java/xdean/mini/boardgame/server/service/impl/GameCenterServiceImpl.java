package xdean.mini.boardgame.server.service.impl;

import java.util.Arrays;
import java.util.Collections;
import java.util.Date;
import java.util.List;
import java.util.Optional;
import java.util.Random;
import java.util.stream.IntStream;

import javax.inject.Inject;

import org.apache.ibatis.session.RowBounds;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.fasterxml.jackson.databind.JsonNode;
import com.google.common.collect.HashMultimap;
import com.google.common.collect.Multimap;

import io.reactivex.Observable;
import io.reactivex.subjects.BehaviorSubject;
import io.reactivex.subjects.Subject;
import xdean.jex.extra.collection.Pair;
import xdean.jex.log.Logable;
import xdean.mini.boardgame.server.model.GameBoard;
import xdean.mini.boardgame.server.model.GlobalConstants;
import xdean.mini.boardgame.server.model.GlobalConstants.AttrKey;
import xdean.mini.boardgame.server.model.GlobalConstants.SocketTopic;
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
import xdean.mini.boardgame.server.service.GameDataService;
import xdean.mini.boardgame.server.service.GameProvider;
import xdean.mini.boardgame.server.service.UserDataService;
import xdean.mini.boardgame.server.socket.AbstractGameSocketProvider;
import xdean.mini.boardgame.server.socket.GameSocketProvider;
import xdean.mini.boardgame.server.socket.WebSocketEvent;
import xdean.mini.boardgame.server.socket.WebSocketSendType;

@Service
public class GameCenterServiceImpl extends AbstractGameSocketProvider implements GameCenterService, GameSocketProvider, Logable {

  @Autowired(required = false)
  List<GameProvider<?>> games = Collections.emptyList();

  private @Inject UserDataService userService;
  private @Inject GameDataService gameMapper;

  private final Object[] locks = IntStream.range(0, 32).mapToObj(i -> new Object()).toArray();

  @Override
  public CreateGameResponse createGame(CreateGameRequest request) {
    Optional<GameProvider<?>> game = findGame(request.getGameName());
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
      GamePlayerEntity player = gameMapper.findPlayer(e.getId());
      if (player.getRoom() != null) {
        return CreateGameResponse.builder()
            .errorCode(GameCenterErrorCode.ALREADY_IN_ROOM)
            .build();
      }
      Integer roomId = generateId();
      GameRoomEntity room = GameRoomEntity.builder()
          .id(roomId)
          .gameName(request.getGameName())
          .createdTime(new Date())
          .playerCount(request.getPlayerCount())
          .roomName(request.getRoomName().isEmpty() ? "Room " + roomId : request.getRoomName())
          .player(player)
          .build();
      GameBoard board = game.get().createGame(room);
      room.setBoard(board);

      player.setRoom(room);
      player.setSeat(0);
      gameMapper.save(room);
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
        Optional<GameRoomEntity> oRoom = gameMapper.findRoom(request.getRoomId());
        if (!oRoom.isPresent()) {
          return JoinGameResponse.builder()
              .errorCode(GameCenterErrorCode.NO_SUCH_ROOM)
              .build();
        }
        GameRoomEntity room = oRoom.get();
        GamePlayerEntity player = gameMapper.findPlayer(user.get().getId());
        if (player.getRoom() != null) {
          return JoinGameResponse.builder()
              .errorCode(GameCenterErrorCode.ALREADY_IN_ROOM)
              .build();
        }
        if (room.getPlayers().size() == room.getPlayerCount()) {
          return JoinGameResponse.builder()
              .errorCode(GameCenterErrorCode.ROOM_FULL)
              .build();
        }
        player.setRoom(room);
        IntStream.range(0, room.getPlayerCount())
            .filter(i -> room.getPlayers().stream().noneMatch(e -> e.getSeat() == i))
            .findFirst()
            .ifPresent(player::setSeat);
        room.addPlayer(player);
        gameMapper.save(room);
        // gamePlayerRepo.save(player);
        int playerId = player.getId();
        sendEvent(playerId, WebSocketEvent.builder()
            .topic(SocketTopic.PLAYER_JOIN)
            .attribute(GlobalConstants.AttrKey.USER_ID, playerId)
            .build());
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
      GamePlayerEntity player = gameMapper.findPlayer(user.get().getId());
      Optional<GameRoomEntity> oRoom = player.getRoom();
      if (!oRoom.isPresent()) {
        return ExitGameResponse.builder()
            .errorCode(GameCenterErrorCode.NOT_IN_ROOM)
            .build();
      }
      GameRoomEntity room = oRoom.get();
      synchronized (getLock(room.getId())) {
        player.setRoom(null);
        gameMapper.save(player);
        room.removePlayer(player);
        if (room.getPlayers().isEmpty()) {
          gameMapper.delete(room);
        } else {
          gameMapper.save(room);
        }
        int playerId = player.getId();
        sendEvent(playerId, WebSocketEvent.builder()
            .topic(SocketTopic.PLAYER_EXIT)
            .attribute(GlobalConstants.AttrKey.USER_ID, playerId)
            .build());
        if (room.getPlayers().isEmpty()) {
          sendEvent(player.getId(), WebSocketEvent.builder()
              .topic(SocketTopic.ROOM_CANCEL)
              .build());
        }
        return ExitGameResponse.builder().build();
      }
    }
  }

  @Override
  public SearchGameResponse searchGame(SearchGameRequest request) {
    Optional<GameProvider<?>> game = findGame(request.getGameName());
    if (!game.isPresent()) {
      return SearchGameResponse.builder()
          .errorCode(GameCenterErrorCode.NO_SUCH_GAME)
          .build();
    }
    List<GameRoomEntity> rooms = gameMapper.findAllByRoomGameName(request.getGameName(),
        new RowBounds(request.getPage() * request.getPageSize(), request.getPageSize()));
    return SearchGameResponse.builder()
        .rooms(rooms)
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
      GamePlayerEntity player = gameMapper.findPlayer(user.get().getId());
      Optional<GameRoomEntity> oRoom = player.getRoom();
      if (!oRoom.isPresent()) {
        return CurrentGameResponse.builder()
            .errorCode(GameCenterErrorCode.NOT_IN_ROOM)
            .build();
      }
      GameRoomEntity room = oRoom.get();
      return CurrentGameResponse.builder().room(room).build();
    }
  }

  /**********
   * SOCKET *
   **********/
  Multimap<Integer, Pair<Integer, Integer>> changeSeatRequests = HashMultimap.create();

  @Override
  protected Subject<WebSocketEvent<?>> createOutputFlow(SocketContext context) {
    return BehaviorSubject.createDefault(WebSocketEvent.builder()
        .type(WebSocketSendType.ALL)
        .topic(SocketTopic.PLAYER_CONNECT)
        .attribute(GlobalConstants.AttrKey.USER_ID, context.getUserId())
        .build());
  }

  @Override
  protected Observable<WebSocketEvent<JsonNode>> processInputFlow(SocketContext context) {
    return context.inputFlow
        .doOnNext(e -> {
          switch (e.getTopic()) {
          case SocketTopic.CHANGE_SEAT_REQUEST:
            changeSeat(context, e);
            break;
          }
        })
        .doOnComplete(() -> context.outputObserver.onNext(WebSocketEvent.builder()
            .topic(SocketTopic.PLAYER_DISCONNECT)
            .attribute(GlobalConstants.AttrKey.USER_ID, context.userId)
            .build()));
  }

  private void changeSeat(SocketContext context, WebSocketEvent<?> e) {
    GameRoomEntity roomEntity = gameMapper.findRoom(context.room.getId()).orElseThrow(IllegalStateException::new);
    int toSeat = ((Number) e.getAttributes().get(AttrKey.TO_SEAT)).intValue();
    synchronized (getLock(context.room.getId())) {
      synchronized (getLock(context.userId)) {
        GamePlayerEntity fromUser = roomEntity.getPlayers().stream().filter(p -> p.getId() == context.userId).findFirst()
            .orElseThrow(IllegalStateException::new);
        int fromSeat = fromUser.getSeat();
        if (fromSeat == toSeat) {
          return;
        }
        Optional<GamePlayerEntity> toUser = roomEntity.getPlayers().stream().filter(p -> p.getSeat() == toSeat).findFirst();
        if (toUser.isPresent()) {
          boolean reverseChange = changeSeatRequests.remove(context.room.getId(), Pair.of(toSeat, fromSeat));
          if (reverseChange) {
            fromUser.setSeat(toSeat);
            toUser.get().setSeat(fromSeat);
            gameMapper.saveAll(Arrays.asList(fromUser, toUser.get()));
          } else {
            Pair<Integer, Integer> changeSeat = Pair.of(fromSeat, toSeat);
            changeSeatRequests.remove(context.room.getId(), changeSeat);
            changeSeatRequests.put(context.room.getId(), changeSeat);
            sendEvent(toUser.get().getId(), WebSocketEvent.builder()
                .type(WebSocketSendType.SELF)
                .topic(SocketTopic.CHANGE_SEAT_REQUEST)
                .attribute(AttrKey.FROM_SEAT, fromSeat)
                .attribute(AttrKey.TO_SEAT, toSeat)
                .build());
            return;
          }
        } else {
          fromUser.setSeat(toSeat);
          gameMapper.save(fromUser);
        }
        sendEvent(context.userId, WebSocketEvent.builder()
            .topic(SocketTopic.CHANGE_SEAT)
            .attribute(AttrKey.FROM_SEAT, fromSeat)
            .attribute(AttrKey.TO_SEAT, toSeat)
            .build());
      }
    }
  }

  private Integer generateId() {
    Random r = new Random();
    Integer id;
    do {
      id = r.nextInt(1000000);
    } while (gameMapper.roomExist(id));
    return id;
  }

  private Optional<GameProvider<?>> findGame(String name) {
    return games.stream().filter(g -> g.name().equals(name)).findFirst();
  }

  private Object getLock(int id) {
    return locks[id % 32];
  }
}
