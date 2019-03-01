package xdean.mini.boardgame.server.wechat.model;

import com.fasterxml.jackson.annotation.JsonAlias;

import lombok.Data;

@Data
public class WechatAuthInfo {
  @JsonAlias("openid")
  String openId;

  @JsonAlias("session_key")
  String sessionKey;

  @JsonAlias("unionid")
  String unionId;

  @JsonAlias("errcode")
  int errorCode;

  @JsonAlias("errmsg")
  String errorMessage;
}
