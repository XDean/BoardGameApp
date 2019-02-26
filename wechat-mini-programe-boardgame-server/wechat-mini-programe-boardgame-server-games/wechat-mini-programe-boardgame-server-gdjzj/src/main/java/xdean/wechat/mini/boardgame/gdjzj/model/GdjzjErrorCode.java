package xdean.wechat.mini.boardgame.gdjzj.model;

import xdean.wechat.mini.boardgame.server.model.exception.MiniBoardgameErrorCode;

public enum GdjzjErrorCode implements MiniBoardgameErrorCode {
  ILLEGAL_PLAYER,
  ILLEGAL_ROLE,
  LAST_PLAYER_IN_TURN,
  PLAYER_ALREADY_DONE,
  ATTACK_SELF

  ;
}
