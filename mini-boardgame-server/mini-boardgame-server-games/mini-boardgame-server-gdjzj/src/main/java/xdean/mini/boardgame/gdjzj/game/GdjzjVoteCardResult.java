package xdean.mini.boardgame.gdjzj.game;

import lombok.Value;

@Value
public class GdjzjVoteCardResult {
  public final int first;
  public final int second;

  public int getScore(GdjzjBoard board) {
    int score = 0;
    if (first != -1) {
      score += board.cards.get(first).real ? 1 : 0;
    }
    if (second != -1) {
      score += board.cards.get(second).real ? 1 : 0;
    }
    return score;
  }
}
