package xdean.wechat.mini.boardgame.gdjzj.game;

import java.util.Arrays;

import xdean.jex.extra.collection.IntList;
import xdean.wechat.mini.boardgame.gdjzj.model.GdjzjErrorCode;
import xdean.wechat.mini.boardgame.server.model.Player;
import xdean.wechat.mini.boardgame.server.model.exception.MiniBoardgameException;

public class GdjzjPlayer {
  static class TurnInfo {
    int order = -1;
    boolean skip = false;
    boolean attack = false;
    int watchCard = -1;
    GdjzjWatchCardResult cardResult;
    int watchCard2 = -1;
    GdjzjWatchCardResult cardResult2;
    int watchPlayer = -1;
    GdjzjWatchPlayerResult playerResult;
  }

  final Player player;
  final int index;
  final GdjzjBoard board;
  final GdjzjRole role;
  TurnInfo[] turnInfos = new TurnInfo[3];

  public GdjzjPlayer(Player player, int index, GdjzjBoard board, GdjzjRole role) {
    this.player = player;
    this.index = index;
    this.board = board;
    this.role = role;
    for (int i = 0; i < 3; i++) {
      turnInfos[i] = new TurnInfo();
    }
  }

  public void selectNextPlayer(int index) {
    checkCurrent();
    IntList leftPlayers = board.getLeftPlayers();
    if (leftPlayers.isEmpty()) {
      throw MiniBoardgameException.builder()
          .message("You are the last player in this turn. Can't select next player.")
          .code(GdjzjErrorCode.LAST_PLAYER_IN_TURN)
          .build();
    }
    if (!leftPlayers.contains(index)) {
      throw MiniBoardgameException.builder()
          .message("Can't select the player as next because it already done in this turn.")
          .code(GdjzjErrorCode.PLAYER_ALREADY_DONE)
          .build();
    }
    board.currentPlayer.set(index);
  }

  public GdjzjWatchCardResult watchCard(int index) {
    checkCurrent();
    GdjzjWatchCardResult result;
    if (isAttacked()) {
      result = GdjzjWatchCardResult.ATTACKED;
    } else if (board.cards.get(index).real) {
      result = GdjzjWatchCardResult.TRUE;
    } else {
      result = GdjzjWatchCardResult.FALSE;
    }
    if (getTurnInfo().watchCard == -1) {
      getTurnInfo().watchCard = index;
      getTurnInfo().cardResult = result;
    } else {
      checkRole(GdjzjRole.XU_YUAN);
      if (getTurnInfo().watchCard2 != -1) {
        throw MiniBoardgameException.builder()
            .code(GdjzjErrorCode.ILLEGAL_STATE)
            .message("The player has watched cards")
            .build();
      }
      getTurnInfo().watchCard2 = index;
      getTurnInfo().cardResult2 = result;
    }
    return result;
  }

  public GdjzjWatchPlayerResult watchPlayer(int index) {
    checkCurrent();
    checkRole(GdjzjRole.FANG_ZHEN);
    GdjzjWatchPlayerResult result;
    if (isAttacked()) {
      result = GdjzjWatchPlayerResult.ATTACKED;
    } else if (board.players.get(index).role.position) {
      result = GdjzjWatchPlayerResult.TRUE;
    } else {
      result = GdjzjWatchPlayerResult.FALSE;
    }
    getTurnInfo().watchPlayer = index;
    getTurnInfo().playerResult = result;
    return result;
  }

  public void attackPlayer(int index) {
    checkCurrent();
    checkRole(GdjzjRole.YAO_BURAN);
    if (index == this.index) {
      throw MiniBoardgameException.builder()
          .code(GdjzjErrorCode.ATTACK_SELF)
          .build();
    }
    board.players.get(index).turnInfos[board.currentTurn.get()].attack = true;
  }

  private boolean isAttacked() {
    return getTurnInfo().attack;
  }

  private TurnInfo getTurnInfo() {
    return turnInfos[board.currentTurn.get()];
  }

  private void checkCurrent() {
    if (board.currentPlayer.get() != index) {
      throw MiniBoardgameException.builder()
          .code(GdjzjErrorCode.ILLEGAL_PLAYER)
          .build();
    }
  }

  private void checkRole(GdjzjRole... roles) {
    if (!Arrays.asList(roles).contains(role)) {
      throw MiniBoardgameException.builder()
          .message("Illegal role to use the skill. Expect: " + Arrays.toString(roles))
          .code(GdjzjErrorCode.ILLEGAL_ROLE)
          .build();
    }
  }
}
