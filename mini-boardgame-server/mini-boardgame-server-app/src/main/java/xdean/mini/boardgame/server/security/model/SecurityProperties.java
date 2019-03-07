package xdean.mini.boardgame.server.security.model;

import lombok.Data;

@Data
public class SecurityProperties {
  String secretKey;
  long expirationTime = 864_000_000;// 10 days
}