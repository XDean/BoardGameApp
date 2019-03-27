package xdean.mini.boardgame.server.security.handler;

import java.io.IOException;
import java.util.Arrays;
import java.util.Optional;

import javax.inject.Inject;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.http.HttpHeaders;
import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import xdean.mini.boardgame.server.handler.LoginSuccessProvider;
import xdean.mini.boardgame.server.handler.LogoutSuccessProvider;
import xdean.mini.boardgame.server.security.TokenAuthProvider;
import xdean.mini.boardgame.server.security.model.AccessToken;

@Component
public class TokenAuthenticationProvider extends OncePerRequestFilter
    implements AuthenticationProvider, LoginSuccessProvider, LogoutSuccessProvider {

  public static final String ACCESS_TOKEN = "access-token";

  @Inject
  TokenAuthProvider provider;

  @Override
  protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain)
      throws ServletException, IOException {
    String token = null;
    Cookie[] cookies = request.getCookies();
    if (cookies != null) {
      Optional<Cookie> cookie = Arrays.stream(cookies)
          .filter(c -> c.getName().equals(ACCESS_TOKEN))
          .findFirst();
      if (cookie.isPresent()) {
        token = cookie.get().getValue();
      }
    }
    if (token == null) {
      token = request.getHeader(HttpHeaders.AUTHORIZATION);
    }
    if (token == null) {
      filterChain.doFilter(request, response);
      return;
    }
    Authentication auth = new AccessToken(token);
    SecurityContextHolder.getContext().setAuthentication(auth);
    filterChain.doFilter(request, response);
  }

  @Override
  public Authentication authenticate(Authentication authentication) throws AuthenticationException {
    String token = ((AccessToken) authentication).getToken();
    return provider.authenticate(token);
  }

  @Override
  public boolean supports(Class<?> authentication) {
    return AccessToken.class.isAssignableFrom(authentication);
  }

  @Override
  public void afterSuccessLogin(HttpServletRequest request, HttpServletResponse response, String username) {
    String token = provider.generateToken(username);
    response.addCookie(new Cookie(ACCESS_TOKEN, token));
    response.addHeader(HttpHeaders.AUTHORIZATION, ACCESS_TOKEN + " " + token);
  }

  @Override
  public void afterSuccessLogout(HttpServletRequest request, HttpServletResponse response) {
  }
}
