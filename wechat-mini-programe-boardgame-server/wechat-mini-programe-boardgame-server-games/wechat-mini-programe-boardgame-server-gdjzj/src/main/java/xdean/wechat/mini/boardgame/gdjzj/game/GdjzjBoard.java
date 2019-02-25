package xdean.wechat.mini.boardgame.gdjzj.game;

import java.util.Comparator;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import com.google.common.collect.ImmutableList;

import javafx.beans.property.IntegerProperty;
import javafx.beans.property.SimpleIntegerProperty;
import xdean.wechat.mini.boardgame.server.model.Player;

public class GdjzjBoard {
  ImmutableList<GdjzjCard> cards;
  ImmutableList<GdjzjPlayer> players;

  IntegerProperty currentPlayer = new SimpleIntegerProperty(this, "currentPlayer");

  public GdjzjBoard(Player[] players) {
    this.cards = createCards();
    List<GdjzjRole> roles = GdjzjRole.getRoles(players.length);
    this.players = ImmutableList.copyOf(IntStream.range(0, players.length)
        .mapToObj(i -> new GdjzjPlayer(players[i], i, this, roles.get(i)))
        .collect(Collectors.toList()));
  }

  private ImmutableList<GdjzjCard> createCards() {
    return ImmutableList.copyOf(IntStream.range(0, 12)
        .mapToObj(i -> new GdjzjCard(i, i / 2 == 0))
        .sorted(Comparator.comparing(c -> c.index % 4 + Math.random()))
        .collect(Collectors.toList()));
  }
}
