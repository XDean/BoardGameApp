package xdean.mini.boardgame.server.security.model;

import java.util.Collection;
import java.util.Collections;

import org.springframework.security.authentication.AbstractAuthenticationToken;
import org.springframework.security.core.GrantedAuthority;

public class AccessToken extends AbstractAuthenticationToken {

  private final String token;

  public AccessToken(String token) {
    this(token, Collections.emptyList());
  }

  public AccessToken(String token, Collection<? extends GrantedAuthority> authorities) {
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
