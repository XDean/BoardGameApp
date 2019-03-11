package xdean.mini.boardgame.server.security.handler;

import java.util.Date;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

import javax.inject.Inject;

import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.CredentialsExpiredException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.User;
import org.springframework.stereotype.Component;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.interfaces.DecodedJWT;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.mybatis.mapper.UserMapper;
import xdean.mini.boardgame.server.security.TokenAuthProvider;
import xdean.mini.boardgame.server.security.model.SecurityProperties;

@Component
public class JwtTokenHandler implements TokenAuthProvider {

  private static final String JWT_TOKEN = "jwt-token";

  private @Inject UserMapper userMapper;
  private @Inject SecurityProperties properties;

  @Override
  public Authentication authenticate(String token) throws AuthenticationException {
    token = token.substring(JWT_TOKEN.length() + 1);
    DecodedJWT verify = JWT.require(Algorithm.HMAC512(properties.getSecretKey().getBytes()))
        .build()
        .verify(token);
    Date expiresAt = verify.getExpiresAt();
    if (expiresAt == null) {
      throw new BadCredentialsException("The token is invalid");
    }
    if (expiresAt.before(new Date())) {
      throw new CredentialsExpiredException("The token is expired");
    }
    String username = verify.getSubject();
    Optional<UserEntity> e = userMapper.findByUsername(username);
    if (e.isPresent()) {
      UserEntity user = e.get();
      List<SimpleGrantedAuthority> authorities = user.getAuthorities().stream()
          .map(a -> new SimpleGrantedAuthority(a)).collect(Collectors.toList());
      return new UsernamePasswordAuthenticationToken(
          new User(user.getUsername(), user.getPassword(), user.isEnabled(), true, true, true, authorities), null, authorities);
    }
    return null;
  }

  @Override
  public String generateToken(String username) {
    return JWT_TOKEN + "-" + JWT.create()
        .withSubject(username)
        .withExpiresAt(new Date(System.currentTimeMillis() + properties.getExpirationTime()))
        .sign(Algorithm.HMAC512(properties.getSecretKey().getBytes()));
  }
}
