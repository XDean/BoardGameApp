package xdean.mini.boardgame.server.security;

import org.springframework.security.core.Authentication;

public interface AuthenticationSuccessProvider {
  void onAuthenticationSuccess(Authentication authentication);
}
