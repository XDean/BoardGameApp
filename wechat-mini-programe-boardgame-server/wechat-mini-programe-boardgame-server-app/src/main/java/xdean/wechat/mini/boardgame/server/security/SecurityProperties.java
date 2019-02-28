package xdean.wechat.mini.boardgame.server.security;

import lombok.Data;

@Data
public class SecurityProperties {
  String secretKey;
  long expirationTime = 864_000_000;// 10 days
  String tokenPrefix = "Bearer ";
}