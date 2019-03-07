package xdean.mini.boardgame.server.security.handler;

import java.io.IOException;
import java.util.Arrays;
import java.util.Date;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

import javax.inject.Inject;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.http.HttpHeaders;
import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.CredentialsExpiredException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.interfaces.DecodedJWT;

import xdean.mini.boardgame.server.handler.LoginSuccessHandler;
import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.security.model.JwtToken;
import xdean.mini.boardgame.server.security.model.SecurityProperties;
import xdean.mini.boardgame.server.service.UserEntityRepo;

@Component
public class JwtTokenHandler extends OncePerRequestFilter implements AuthenticationProvider, LoginSuccessHandler {

  public static final String JWT_TOKEN = "jwt-token";

  @Inject
  UserEntityRepo userRepo;

  @Inject
  SecurityProperties properties;

  @Override
  protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain)
      throws ServletException, IOException {
    String token = null;
    Optional<Cookie> cookie = Arrays.stream(request.getCookies())
        .filter(c -> c.getName().equals(JWT_TOKEN))
        .findFirst();
    if (cookie.isPresent()) {
      token = cookie.get().getValue();
    }
    if (token == null) {
      String header = request.getHeader(HttpHeaders.AUTHORIZATION);
      if (header != null && header.startsWith(JWT_TOKEN)) {
        token = header.substring(JWT_TOKEN.length() + 1);
      }
    }
    if (token == null) {
      filterChain.doFilter(request, response);
      return;
    }

    Authentication auth = new JwtToken(token);
    SecurityContextHolder.getContext().setAuthentication(auth);
    filterChain.doFilter(request, response);
  }

  @Override
  public Authentication authenticate(Authentication authentication) throws AuthenticationException {
    JwtToken token = (JwtToken) authentication;
    DecodedJWT verify = JWT.require(Algorithm.HMAC512(properties.getSecretKey().getBytes()))
        .build()
        .verify(token.getToken());
    Date expiresAt = verify.getExpiresAt();
    if (expiresAt == null) {
      throw new BadCredentialsException("The token is invalid");
    }
    if (expiresAt.before(new Date())) {
      throw new CredentialsExpiredException("The token is expired");
    }
    String username = verify.getSubject();
    Optional<UserEntity> e = userRepo.findByUsername(username);
    if (e.isPresent()) {
      UserEntity user = e.get();
      List<SimpleGrantedAuthority> authorities = user.getAuthorities().stream()
          .map(a -> new SimpleGrantedAuthority(a.getAuthority())).collect(Collectors.toList());
      return new UsernamePasswordAuthenticationToken(
          new User(user.getUsername(), user.getPassword(), user.isEnabled(), true, true, true, authorities), null, authorities);
    }
    return null;
  }

  @Override
  public boolean supports(Class<?> authentication) {
    return JwtToken.class.isAssignableFrom(authentication);
  }

  @Override
  public void afterSuccess(HttpServletRequest request, HttpServletResponse response, String username) {
    String token = JWT.create()
        .withSubject(username)
        .withExpiresAt(new Date(System.currentTimeMillis() + properties.getExpirationTime()))
        .sign(Algorithm.HMAC512(properties.getSecretKey().getBytes()));
    response.addCookie(new Cookie(JWT_TOKEN, token));
    response.addHeader(HttpHeaders.AUTHORIZATION, JWT_TOKEN + " " + token);
  }
}
