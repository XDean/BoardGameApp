package xdean.mini.boardgame.server.wechat.handler;

import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.core.AuthenticationException;
import org.springframework.stereotype.Component;

import xdean.mini.boardgame.server.security.OpenIdAuthProvider;

@Component
public class WechatOpenIdProvder implements OpenIdAuthProvider {

  @Override
  public String name() {
    return "wechat";
  }

  @Override
  public String attemptAuthentication(String token) throws AuthenticationException {
    if (token.equals("abc")) {
      return "abc";
    } else if (token.length() == 1) {
      return null;
    } else {
      throw new BadCredentialsException("bad");
    }
  }
}
