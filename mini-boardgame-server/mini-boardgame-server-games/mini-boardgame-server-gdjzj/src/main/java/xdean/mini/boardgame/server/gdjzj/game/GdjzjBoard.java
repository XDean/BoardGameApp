package xdean.mini.boardgame.server.gdjzj.game;

import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;
import java.util.Map.Entry;
import java.util.Optional;
import java.util.concurrent.ThreadLocalRandom;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.google.common.collect.ImmutableList;

import xdean.jex.extra.collection.IntList;
import xdean.mini.boardgame.server.gdjzj.model.GdjzjErrorCode;
import xdean.mini.boardgame.server.model.GameBoard;
import xdean.mini.boardgame.server.model.entity.GamePlayerEntity;
import xdean.mini.boardgame.server.model.entity.GameRoomEntity;
import xdean.mini.boardgame.server.model.exception.MiniBoardgameException;

public class GdjzjBoard extends GameBoard {
  List<GdjzjCard> cards = Collections.emptyList();
  List<GdjzjPlayer> players = Collections.emptyList();

  int currentPlayer = 0;
  int currentTurn = 0;

  @JsonCreator
  public GdjzjBoard(@JsonProperty("room") GameRoomEntity room) {
    super(room);
  }

  @Override
  public void start(GamePlayerEntity[] players) {
    checkState(WAITING);
    Arrays.sort(players, Comparator.comparing(p -> p.getSeat()));
    this.cards = createCards();
    List<GdjzjRole> roles = GdjzjRole.getRoles(players.length);
    this.players = ImmutableList.copyOf(IntStream.range(0, players.length)
        .mapToObj(i -> new GdjzjPlayer(players[i].getUserId(), i, this, roles.get(i)))
        .collect(Collectors.toList()));
    this.currentPlayer = ThreadLocalRandom.current().nextInt(players.length);
    this.currentTurn = 0;
    state = START;
  }

  public void nextTurn() {
    if (currentTurn < 2) {
      currentTurn++;
    }
  }

  public GdjzjVoteCardResult getVoteCardResult(int turn) {
    assertTurn(turn);
    List<Integer> result = players.stream()
        .map(p -> p.turnInfos[turn])
        .filter(t -> t.vote)
        .flatMapToInt(t -> IntStream.of(t.vote1, t.vote2))
        .filter(i -> i != -1)
        .boxed()
        .collect(Collectors.groupingBy(t -> t))
        .entrySet()
        .stream()
        .sorted(Comparator.<Entry<Integer, List<Integer>>, Integer> comparing(e -> -e.getValue().size())
            .thenComparing(e -> e.getKey()))
        .map(e -> e.getKey())
        .limit(2L)
        .collect(Collectors.toList());
    return new GdjzjVoteCardResult(result.size() > 0 ? result.get(0) : -1, result.size() > 1 ? result.get(1) : -1);
  }

  public int getScore() {
    int score = 0;
    score += getVoteCardResult(0).getScore(this);
    score += getVoteCardResult(1).getScore(this);
    score += getVoteCardResult(2).getScore(this);
    if (score == 6) {
      return score;
    }

    Optional<GdjzjPlayer> boss = getPlayer(GdjzjRole.LAO_CHAOFENG);
    int bossIndex = boss.map(p -> p.index).orElse(-1);
    if (!boss
        .map(p -> p.vote)
        .filter(i -> i != -1)
        .map(i -> players.get(i).role == GdjzjRole.XU_YUAN)
        .orElse(false)) {
      score += 2;
    }
    if (!getPlayer(GdjzjRole.YAO_BURAN)
        .map(p -> p.vote)
        .filter(i -> i != -1)
        .map(i -> players.get(i).role == GdjzjRole.FANG_ZHEN)
        .orElse(false)) {
      score += 1;
    }
    long playerCount = players.stream()
        .filter(p -> p.role.position)
        .count();
    long rightCount = players.stream()
        .filter(p -> p.role.position)
        .map(p -> p.vote)
        .filter(i -> i != -1)
        .filter(i -> i == bossIndex)
        .count();
    if (rightCount > playerCount / 2) {
      score += 1;
    }
    return score;
  }

  public IntList getLeftPlayers() {
    return IntList.create(players.stream().filter(p -> p.turnInfos[currentTurn].order < 0).mapToInt(p -> p.index).toArray());
  }

  public Optional<GdjzjPlayer> getPlayer(GdjzjRole role) {
    return players.stream().filter(p -> p.role == role).findFirst();
  }

  private ImmutableList<GdjzjCard> createCards() {
    List<GdjzjCard> collect = IntStream.range(0, 12)
        .mapToObj(i -> new GdjzjCard(i, i / 2 == 0))
        .sorted(Comparator.comparing(c -> c.index % 4 + Math.random()))
        .collect(Collectors.toList());
    return ImmutableList.copyOf(IntStream.range(0, 12)
        .mapToObj(i -> new GdjzjCard(i, collect.get(i).real))
        .collect(Collectors.toList()));
  }

  public static void assertTurn(int turn) {
    if (turn < 0 || turn > 2) {
      throw MiniBoardgameException.builder()
          .code(GdjzjErrorCode.ILLEGAL_TURN)
          .build();
    }
  }
}
