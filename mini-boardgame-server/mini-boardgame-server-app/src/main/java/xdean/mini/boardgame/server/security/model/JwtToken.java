package xdean.mini.boardgame.server.security.model;

import java.util.Collection;
import java.util.Collections;

import org.springframework.security.authentication.AbstractAuthenticationToken;
import org.springframework.security.core.GrantedAuthority;

public class JwtToken extends AbstractAuthenticationToken {

  private final String token;

  public JwtToken(String token) {
    this(token, Collections.emptyList());
  }

  public JwtToken(String token, Collection<? extends GrantedAuthority> authorities) {
    super(authorities);
    this.token = token;
  }

  @Override
  public Object getPrincipal() {
    return token;
  }

  @Override
  public Object getCredentials() {
    return null;
  }

  public String getToken() {
    return token;
  }
}
