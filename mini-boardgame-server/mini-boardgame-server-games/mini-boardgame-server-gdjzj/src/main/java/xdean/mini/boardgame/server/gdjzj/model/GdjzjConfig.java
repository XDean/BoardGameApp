package xdean.mini.boardgame.server.gdjzj.model;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import lombok.ToString;
import xdean.mini.boardgame.server.model.GameConfig;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@ToString(callSuper = false)
@EqualsAndHashCode(callSuper = false)
public class GdjzjConfig extends GameConfig {
  int playerCount;
}
