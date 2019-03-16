package xdean.mini.boardgame.server.security;

import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;

public interface AuthenticationFailProvider {
  void onAuthenticationFail(AuthenticationException exception, Authentication authentication);
}
