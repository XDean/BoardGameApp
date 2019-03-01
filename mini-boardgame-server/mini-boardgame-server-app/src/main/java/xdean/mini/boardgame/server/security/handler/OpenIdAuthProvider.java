package xdean.mini.boardgame.server.security.handler;

import org.springframework.security.core.AuthenticationException;

public interface OpenIdAuthProvider {
  String name();

  String attemptAuthentication(String token) throws AuthenticationException;
}
