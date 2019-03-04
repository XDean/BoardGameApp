package xdean.mini.boardgame.server.gdjzj.model;

import xdean.mini.boardgame.server.model.exception.MiniBoardgameErrorCode;

public enum GdjzjErrorCode implements MiniBoardgameErrorCode {
  ILLEGAL_STATE,
  ILLEGAL_TURN,
  ILLEGAL_PLAYER,
  ILLEGAL_CARD,
  ILLEGAL_ROLE,
  LAST_PLAYER_IN_TURN,
  PLAYER_ALREADY_DONE,
  ATTACK_SELF

  ;
}
