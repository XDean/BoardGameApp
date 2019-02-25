package xdean.wechat.mini.boardgame.gdjzj.game;

import xdean.wechat.mini.boardgame.gdjzj.model.GdjzjErrorCode;
import xdean.wechat.mini.boardgame.server.model.Player;
import xdean.wechat.mini.boardgame.server.model.exception.MiniBoardgameException;

public class GdjzjPlayer {
  final Player player;
  final int index;
  final GdjzjBoard board;
  final GdjzjRole role;
  int turn;

  public GdjzjPlayer(Player player, int index, GdjzjBoard board, GdjzjRole role) {
    this.player = player;
    this.index = index;
    this.board = board;
    this.role = role;
  }

  public GdjzjWatchCardResult watchCard(int index) {
    checkCurrent();
    return board.cards.get(index).real ? GdjzjWatchCardResult.TRUE : GdjzjWatchCardResult.FALSE;
  }

  public void selectNextPlayer(int index) {
    checkCurrent();
    board.currentPlayer.set(index);
  }

  private void checkCurrent() {
    if (board.currentPlayer.get() != index) {
      throw MiniBoardgameException.builder()
          .code(GdjzjErrorCode.ILLEGAL_PLAYER)
          .build();
    }
  }
}
