package xdean.wechat.mini.boardgame.gdjzj.game;

import java.util.Arrays;

import xdean.jex.extra.collection.IntList;
import xdean.wechat.mini.boardgame.gdjzj.model.GdjzjErrorCode;
import xdean.wechat.mini.boardgame.server.model.Player;
import xdean.wechat.mini.boardgame.server.model.exception.MiniBoardgameException;

public class GdjzjPlayer {
  public static class TurnInfo {
    int order = -1;
    boolean skip = false;
    boolean attack = false;
    int checkCard = -1;
    GdjzjCheckCardResult cardResult;
    int checkCard2 = -1;
    GdjzjCheckCardResult cardResult2;
    int checkPlayer = -1;
    GdjzjCheckPlayerResult playerResult;
    boolean reverseCard = false;
  }

  final Player player;
  final int index;
  final GdjzjBoard board;
  final GdjzjRole role;
  final TurnInfo[] turnInfos = new TurnInfo[3];

  public GdjzjPlayer(Player player, int index, GdjzjBoard board, GdjzjRole role) {
    this.player = player;
    this.index = index;
    this.board = board;
    this.role = role;
    for (int i = 0; i < 3; i++) {
      turnInfos[i] = new TurnInfo();
      if (role == GdjzjRole.FANG_ZHEN) {
        turnInfos[i].skip = true;
      }
    }
    if (role == GdjzjRole.HUANG_YANYAN || role == GdjzjRole.MUHU_JIANAI) {
      turnInfos[(int) (Math.random() * 3)].skip = true;
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

  public GdjzjCheckCardResult checkCard(int index) {
    checkCurrent();
    if (index / 4 != board.currentTurn.get()) {
      throw MiniBoardgameException.builder()
          .code(GdjzjErrorCode.ILLEGAL_CARD)
          .message("This card can't be checked in this turn")
          .build();
    }
    GdjzjCheckCardResult result;
    GdjzjCard card = board.cards.get(index);
    if (getTurnInfo().skip) {
      result = GdjzjCheckCardResult.SKIPPED;
    } else if (getTurnInfo().attack) {
      result = GdjzjCheckCardResult.ATTACKED;
    } else if ((role.position && role != GdjzjRole.JI_YUNFU) ? card.real ^ card.reverse : card.real) {
      result = GdjzjCheckCardResult.TRUE;
    } else {
      result = GdjzjCheckCardResult.FALSE;
    }
    if (getTurnInfo().checkCard == -1) {
      getTurnInfo().checkCard = index;
      getTurnInfo().cardResult = result;
    } else {
      checkRole(GdjzjRole.XU_YUAN);
      if (getTurnInfo().checkCard2 != -1) {
        throw MiniBoardgameException.builder()
            .code(GdjzjErrorCode.ILLEGAL_STATE)
            .message("The player has checked cards")
            .build();
      }
      getTurnInfo().checkCard2 = index;
      getTurnInfo().cardResult2 = result;
    }
    return result;
  }

  public GdjzjCheckPlayerResult checkPlayer(int index) {
    checkCurrent();
    checkRole(GdjzjRole.FANG_ZHEN);
    GdjzjCheckPlayerResult result;
    if (getTurnInfo().attack) {
      result = GdjzjCheckPlayerResult.ATTACKED;
    } else if (board.players.get(index).role.position) {
      result = GdjzjCheckPlayerResult.TRUE;
    } else {
      result = GdjzjCheckPlayerResult.FALSE;
    }
    getTurnInfo().checkPlayer = index;
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
    board.players.get(index).attacked();
  }

  private void attacked() {
    int turn = board.currentTurn.get();
    if (turnInfos[turn].order != -1) {
      turn++;
    }
    if (turn < 3) {
      if (role == GdjzjRole.JI_YUNFU) {
        for (int i = turn; i < 3; i++) {
          turnInfos[i].attack = true;
        }
      } else {
        turnInfos[turn].attack = true;
      }
    }
    if (role == GdjzjRole.FANG_ZHEN) {
      board.getPlayer(GdjzjRole.XU_YUAN).ifPresent(p -> p.attacked());
    }
  }

  public void reverseCard() {
    checkCurrent();
    checkRole(GdjzjRole.LAO_CHAOFENG);
    if (!getTurnInfo().reverseCard) {
      getTurnInfo().reverseCard = true;
      board.cards.forEach(c -> c.reverse = true);
    }
  }

  public void voteCard(int a, int b) {

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
