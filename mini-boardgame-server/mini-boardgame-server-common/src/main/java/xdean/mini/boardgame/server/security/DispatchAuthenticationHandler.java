package xdean.mini.boardgame.server.security;

import java.util.Collections;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationEventPublisher;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.stereotype.Component;

@Component
public class DispatchAuthenticationHandler implements AuthenticationEventPublisher {
  @Autowired(required = false)
  List<AuthenticationSuccessProvider> successProviders = Collections.emptyList();

  @Autowired(required = false)
  List<AuthenticationFailProvider> failProviders = Collections.emptyList();

  @Override
  public void publishAuthenticationSuccess(Authentication authentication) {
    successProviders.forEach(p -> p.onAuthenticationSuccess(authentication));
  }

  @Override
  public void publishAuthenticationFailure(AuthenticationException exception, Authentication authentication) {
    failProviders.forEach(p -> p.onAuthenticationFail(exception, authentication));
  }
}
