package xdean.mini.boardgame.server.service.impl;

import java.util.Arrays;
import java.util.Collections;
import java.util.Date;
import java.util.List;
import java.util.Optional;
import java.util.Random;
import java.util.stream.IntStream;

import javax.inject.Inject;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.common.collect.HashMultimap;
import com.google.common.collect.Multimap;

import io.reactivex.Observable;
import io.reactivex.subjects.BehaviorSubject;
import io.reactivex.subjects.Subject;
import xdean.jex.extra.collection.Pair;
import xdean.jex.log.Logable;
import xdean.mini.boardgame.server.model.CommonConstants;
import xdean.mini.boardgame.server.model.CommonConstants.AttrKey;
import xdean.mini.boardgame.server.model.CommonConstants.SocketTopic;
import xdean.mini.boardgame.server.model.GameBoard;
import xdean.mini.boardgame.server.model.GameBoard.State;
import xdean.mini.boardgame.server.model.GameConfig;
import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.exception.MiniBoardgameException;
import xdean.mini.boardgame.server.model.param.CreateGameRequest;
import xdean.mini.boardgame.server.model.param.CreateGameResponse;
import xdean.mini.boardgame.server.model.param.CurrentGameResponse;
import xdean.mini.boardgame.server.model.param.ExitGameRequest;
import xdean.mini.boardgame.server.model.param.ExitGameResponse;
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
import xdean.mini.boardgame.server.socket.WebSocketIllegalArgumentException;
import xdean.mini.boardgame.server.socket.WebSocketSendType;

@Service
public class GameCenterServiceImpl extends AbstractGameSocketProvider implements GameCenterService, GameSocketProvider, Logable {

  @Autowired(required = false)
  List<GameProvider<?, ?>> games = Collections.emptyList();

  private @Inject UserDataService userService;
  private @Inject GameDataService gameMapper;
  private @Inject ObjectMapper objectMapper;

  private final Object[] locks = IntStream.range(0, 32).mapToObj(i -> new Object()).toArray();

  @Override
  public CreateGameResponse createGame(CreateGameRequest request) {
    Optional<UserEntity> userOptional = userService.getCurrentUser();
    if (!userOptional.isPresent()) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.UNAUTHORIZED)
          .message("No authorized user")
          .build();
    }
    Optional<GameProvider<?, ?>> gameOptional = findGame(request.getGameName());
    if (!gameOptional.isPresent()) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.NOT_FOUND)
          .message("No such game")
          .build();
    }
    GameProvider<?, ?> game = gameOptional.get();
    UserEntity user = userOptional.get();
    synchronized (getLock(user.getId())) {
      GamePlayerEntity player = gameMapper.findPlayer(user.getId());
      if (player.getRoom().isPresent()) {
        throw MiniBoardgameException.builder()
            .code(HttpStatus.BAD_REQUEST)
            .message("Already in a room")
            .build();
      }
      Integer roomId = generateId();
      GameRoomEntity room = GameRoomEntity.builder()
          .id(roomId)
          .gameName(request.getGameName())
          .createdTime(new Date())
          // .playerCount(request.getPlayerCount())
          .roomName(request.getRoomName().isEmpty() ? "Room " + roomId : request.getRoomName())
          .player(player)
          .build();
      config(request, game, room);
      GameBoard board = game.createGame(room);
      board.setRoom(room);
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

  private <C extends GameConfig> void config(CreateGameRequest request, GameProvider<?, C> game, GameRoomEntity room) {
    C config = objectMapper.convertValue(request.getGameConfig(), game.configClass());
    game.configRoom(room, config);
  }

  @Override
  public JoinGameResponse joinGame(JoinGameRequest request) {
    Optional<UserEntity> user = userService.getCurrentUser();
    if (!user.isPresent()) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.UNAUTHORIZED)
          .message("No authorized user")
          .build();
    }
    synchronized (getLock(user.get().getId())) {
      synchronized (getLock(request.getRoomId())) {
        Optional<GameRoomEntity> oRoom = gameMapper.findRoom(request.getRoomId());
        if (!oRoom.isPresent()) {
          throw MiniBoardgameException.builder()
              .code(HttpStatus.NOT_FOUND)
              .message("No such room")
              .build();
        }
        GameRoomEntity room = oRoom.get();
        GamePlayerEntity player = gameMapper.findPlayer(user.get().getId());
        if (player.getRoom().isPresent()) {
          throw MiniBoardgameException.builder()
              .code(HttpStatus.BAD_REQUEST)
              .message("Already in a game room")
              .build();
        }
        if (room.getPlayers().size() == room.getPlayerCount()) {
          throw MiniBoardgameException.builder()
              .code(HttpStatus.BAD_REQUEST)
              .message("The room is full")
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
        sendEvent(room.getId(), playerId, WebSocketEvent.builder()
            .topic(SocketTopic.PLAYER_JOIN)
            .attribute(AttrKey.USER_ID, playerId)
            .attribute(AttrKey.SEAT, player.getSeat())
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
      throw MiniBoardgameException.builder()
          .code(HttpStatus.UNAUTHORIZED)
          .message("No authorized user")
          .build();
    }
    synchronized (getLock(user.get().getId())) {
      GamePlayerEntity player = gameMapper.findPlayer(user.get().getId());
      Optional<GameRoomEntity> oRoom = player.getRoom();
      if (!oRoom.isPresent()) {
        throw MiniBoardgameException.builder()
            .code(HttpStatus.BAD_REQUEST)
            .message("Not in game room")
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
        sendEvent(room.getId(), playerId, WebSocketEvent.builder()
            .topic(SocketTopic.PLAYER_EXIT)
            .attribute(CommonConstants.AttrKey.USER_ID, playerId)
            .build());
        if (room.getPlayers().isEmpty()) {
          sendEvent(room.getId(), player.getId(), WebSocketEvent.builder()
              .topic(SocketTopic.ROOM_CANCEL)
              .build());
        }
        return ExitGameResponse.builder().build();
      }
    }
  }

  @Override
  public SearchGameResponse searchGame(SearchGameRequest request) {
    return SearchGameResponse.builder()
        .rooms(gameMapper.searchGame(request))
        .build();
  }

  @Override
  public CurrentGameResponse currentGame() {
    Optional<UserEntity> user = userService.getCurrentUser();
    if (!user.isPresent()) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.UNAUTHORIZED)
          .message("No authorized user")
          .build();
    }
    synchronized (getLock(user.get().getId())) {
      GamePlayerEntity player = gameMapper.findPlayer(user.get().getId());
      Optional<GameRoomEntity> oRoom = player.getRoom();
      if (!oRoom.isPresent()) {
        throw MiniBoardgameException.builder()
            .code(HttpStatus.BAD_REQUEST)
            .message("Not in game room")
            .build();
      }
      GameRoomEntity room = oRoom.get();
      return CurrentGameResponse.builder().playerId(player.getId()).room(room).build();
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
        .attribute(CommonConstants.AttrKey.USER_ID, context.getUserId())
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
          case SocketTopic.PLAYER_READY:
            playerReady(context, e);
            break;
          }
        })
        .doOnComplete(() -> context.outputObserver.onNext(WebSocketEvent.builder()
            .topic(SocketTopic.PLAYER_DISCONNECT)
            .attribute(CommonConstants.AttrKey.USER_ID, context.userId)
            .build()));
  }

  private void changeSeat(SocketContext context, WebSocketEvent<?> e) {
    synchronized (getLock(context.room.getId())) {
      synchronized (getLock(context.userId)) {
        GameRoomEntity roomEntity = gameMapper.findRoom(context.room.getId())
            .orElseThrow(() -> WebSocketIllegalArgumentException.builder().event(e).message("Room canceled").build());
        int toSeat = ((Number) e.getAttributes().get(AttrKey.TO_SEAT)).intValue();
        GamePlayerEntity fromUser = roomEntity.getPlayers().stream().filter(p -> p.getId() == context.userId)
            .findFirst().orElseThrow(() -> WebSocketIllegalArgumentException.builder().event(e)
                .message("Only player in the room can change seat").build());
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
            sendEvent(context.room.getId(), toUser.get().getId(), WebSocketEvent.builder()
                .id(e.getId())
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
        sendEvent(context.room.getId(), context.userId, WebSocketEvent.builder()
            .id(e.getId())
            .topic(SocketTopic.CHANGE_SEAT)
            .attribute(AttrKey.FROM_SEAT, fromSeat)
            .attribute(AttrKey.TO_SEAT, toSeat)
            .build());
      }
    }
  }

  private void playerReady(SocketContext context, WebSocketEvent<JsonNode> e) {
    int roomId = context.room.getId();
    synchronized (getLock(roomId)) {
      synchronized (getLock(context.userId)) {
        GameRoomEntity roomEntity = gameMapper.findRoom(roomId)
            .orElseThrow(() -> WebSocketIllegalArgumentException.builder().event(e).message("Room canceled").build());
        roomEntity.getBoard().checkState(State.WAITING);
        GamePlayerEntity user = roomEntity.getPlayers().stream().filter(p -> p.getId() == context.userId).findFirst()
            .orElseThrow(() -> WebSocketIllegalArgumentException.builder().event(e)
                .message("Only player in the room can be ready").build());
        boolean ready = e.getPayload().booleanValue();
        if (user.isReady() == ready) {
          return;
        }
        user.setReady(ready);
        gameMapper.save(user);
        sendEvent(roomId, context.userId, WebSocketEvent.builder()
            .topic(SocketTopic.PLAYER_READY)
            .attribute(AttrKey.USER_ID, context.userId)
            .payload(ready)
            .build());
        if (roomEntity.getPlayers().stream().filter(p -> p.isReady()).count() == roomEntity.getPlayerCount()) {
          roomEntity.getBoard().setState(State.PLAYING);
          sendEvent(roomId, context.userId, WebSocketEvent.builder()
              .topic(SocketTopic.GAME_START)
              .build());
        }
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

  private Optional<GameProvider<?, ?>> findGame(String name) {
    return games.stream().filter(g -> g.name().equals(name)).findFirst();
  }

  private Object getLock(int id) {
    return locks[id % 32];
  }
}
