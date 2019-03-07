package xdean.mini.boardgame.server.security;

import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;

public interface TokenAuthProvider {
  Authentication authenticate(String token) throws AuthenticationException;

  String generateToken(String username);
}
