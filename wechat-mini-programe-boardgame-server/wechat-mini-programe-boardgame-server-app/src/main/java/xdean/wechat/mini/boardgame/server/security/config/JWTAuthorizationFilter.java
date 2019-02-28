package xdean.wechat.mini.boardgame.server.security.config;

import java.io.IOException;
import java.util.ArrayList;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.http.HttpHeaders;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;

import xdean.wechat.mini.boardgame.server.security.SecurityProperties;

public class JWTAuthorizationFilter extends BasicAuthenticationFilter {
  SecurityProperties properties;

  public JWTAuthorizationFilter(AuthenticationManager authManager, SecurityProperties properties) {
    super(authManager);
    this.properties = properties;
  }

  @Override
  protected void doFilterInternal(HttpServletRequest req, HttpServletResponse res, FilterChain chain)
      throws IOException, ServletException {
    String header = req.getHeader(HttpHeaders.AUTHORIZATION);

    if (header == null || !header.startsWith(properties.getTokenPrefix())) {
      chain.doFilter(req, res);
      return;
    }

    UsernamePasswordAuthenticationToken authentication = getAuthentication(header);

    SecurityContextHolder.getContext().setAuthentication(authentication);
    chain.doFilter(req, res);
  }

  private UsernamePasswordAuthenticationToken getAuthentication(String token) {
    // parse the token.
    String user = JWT.require(Algorithm.HMAC512(properties.getSecretKey().getBytes()))
        .build()
        .verify(token.replace(properties.getTokenPrefix(), ""))
        .getSubject();
    if (user != null) {
      return new UsernamePasswordAuthenticationToken(user, null, new ArrayList<>());
    }
    return null;
  }
}