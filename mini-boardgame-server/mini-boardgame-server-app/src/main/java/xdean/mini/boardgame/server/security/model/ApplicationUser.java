package xdean.mini.boardgame.server.security.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Value;

@Value
@Builder
@AllArgsConstructor
public class ApplicationUser {
  private String username;
  private String password;
}