package xdean.mini.boardgame.server.gdjzj.game;

import java.util.Arrays;
import java.util.Comparator;
import java.util.List;
import java.util.stream.Collectors;

public enum GdjzjRole {
  XU_YUAN(true, 6),
  FANG_ZHEN(true, 6),
  JI_YUNFU(true, 8),
  HUANG_YANYAN(true, 6),
  MUHU_JIANAI(true, 6),
  LAO_CHAOFENG(false, 6),
  YAO_BURAN(false, 6),
  ZHENG_GUOQU(false, 7);

  public final boolean position;
  public final int limit;

  private GdjzjRole(boolean position, int limit) {
    this.position = position;
    this.limit = limit;
  }

  public static List<GdjzjRole> getRoles(int count) {
    return Arrays.stream(GdjzjRole.values())
        .filter(r -> r.limit <= count)
        .sorted(Comparator.comparing(s -> Math.random()))
        .collect(Collectors.toList());
  }
}
