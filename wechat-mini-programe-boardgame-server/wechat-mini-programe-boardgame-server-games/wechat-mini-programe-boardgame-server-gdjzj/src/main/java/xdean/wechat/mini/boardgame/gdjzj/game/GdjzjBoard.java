package xdean.wechat.mini.boardgame.gdjzj.game;

import javafx.beans.property.IntegerProperty;
import javafx.beans.property.SimpleIntegerProperty;

public class GdjzjBoard {
  GdjzjCard[] cards;
  GdjzjPlayer[] players;

  IntegerProperty currentPlayer = new SimpleIntegerProperty(this, "currentPlayer");

  public GdjzjBoard() {

  }
}
