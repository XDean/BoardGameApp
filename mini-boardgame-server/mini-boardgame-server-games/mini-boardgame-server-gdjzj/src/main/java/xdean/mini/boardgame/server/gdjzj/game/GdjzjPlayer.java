package xdean.mini.boardgame.server.gdjzj.game;

import java.util.Arrays;

import org.springframework.http.HttpStatus;

import xdean.jex.extra.collection.IntList;
import xdean.mini.boardgame.server.model.exception.MiniBoardgameException;

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

    boolean vote;
    int vote1;
    int vote2;
  }

  final int id;
  final int index;
  final GdjzjBoard board;
  final GdjzjRole role;
  final TurnInfo[] turnInfos = new TurnInfo[3];
  int vote;

  public GdjzjPlayer(int id, int index, GdjzjBoard board, GdjzjRole role) {
    this.id = id;
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
    assertCurrentPlayer();
    IntList leftPlayers = board.getLeftPlayers();
    if (leftPlayers.isEmpty()) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.BAD_REQUEST)
          .message("You are the last player in this turn. Can't select next player.")
          .build();
    }
    if (!leftPlayers.contains(index)) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.BAD_REQUEST)
          .message("Can't select the player as next because it already done in this turn.")
          .build();
    }
    board.currentPlayer = index;
  }

  public GdjzjCheckCardResult checkCard(int index) {
    assertCurrentPlayer();
    assertCard(index);
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
      assertRole(GdjzjRole.XU_YUAN);
      if (getTurnInfo().checkCard2 != -1) {
        throw MiniBoardgameException.builder()
            .code(HttpStatus.BAD_REQUEST)
            .message("The player has checked cards")
            .build();
      }
      getTurnInfo().checkCard2 = index;
      getTurnInfo().cardResult2 = result;
    }
    return result;
  }

  public GdjzjCheckPlayerResult checkPlayer(int index) {
    assertCurrentPlayer();
    assertRole(GdjzjRole.FANG_ZHEN);
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
    assertCurrentPlayer();
    assertRole(GdjzjRole.YAO_BURAN);
    if (index == this.index) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.BAD_REQUEST)
          .message("You can't attack yourself")
          .build();
    }
    board.players.get(index).attacked();
  }

  private void attacked() {
    int turn = board.currentTurn;
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
    assertCurrentPlayer();
    assertRole(GdjzjRole.LAO_CHAOFENG);
    if (!getTurnInfo().reverseCard) {
      getTurnInfo().reverseCard = true;
      board.cards.forEach(c -> c.reverse = true);
    }
  }

  public void voteCard(int a, int b) {
    if (a != -1) {
      assertCard(a);
    }
    if (b != -1) {
      assertCard(b);
    }
    TurnInfo turnInfo = getTurnInfo();
    if (turnInfo.vote) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.BAD_REQUEST)
          .message("You have voted card this turn")
          .build();
    }
    turnInfo.vote = true;
    turnInfo.vote1 = a;
    turnInfo.vote2 = b;
  }

  public void votePlayer(int index) {
    assertPlayer(index);
    if (vote != -1) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.BAD_REQUEST)
          .message("You have voted player")
          .build();
    }
    vote = index;
  }

  private TurnInfo getTurnInfo() {
    return turnInfos[board.currentTurn];
  }

  private void assertCurrentPlayer() {
    if (board.currentPlayer != index) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.BAD_REQUEST)
          .message("You are not the active player")
          .build();
    }
  }

  private void assertCard(int index) {
    if (index / 4 != board.currentTurn) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.BAD_REQUEST)
          .message("This card can't be used in this turn")
          .build();
    }
  }

  private void assertPlayer(int index) {
    if (index < 0 || index >= board.players.size()) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.BAD_REQUEST)
          .message("Illegal player index")
          .build();
    }
  }

  private void assertRole(GdjzjRole... roles) {
    if (!Arrays.asList(roles).contains(role)) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.BAD_REQUEST)
          .message("Illegal role to use the skill. Expect: " + Arrays.toString(roles))
          .build();
    }
  }
}
