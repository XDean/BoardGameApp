package xdean.mini.boardgame.server.wechat.handler;

import javax.inject.Inject;

import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.core.AuthenticationException;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

import com.google.common.collect.ImmutableMap;

import xdean.jex.log.Logable;
import xdean.mini.boardgame.server.security.OpenIdAuthProvider;
import xdean.mini.boardgame.server.wechat.config.WechatProperties;
import xdean.mini.boardgame.server.wechat.model.WechatAuthInfo;

@Component
public class WechatOpenIdProvder implements OpenIdAuthProvider, Logable {

  @Inject
  RestTemplate restTemplate;

  @Inject
  WechatProperties wechatProperties;

  @Override
  public String name() {
    return "wechat-mbg";
  }

  @Override
  public String attemptAuthentication(String token) throws AuthenticationException {
    WechatAuthInfo info;
    do {
      info = restTemplate.getForObject(
          wechatProperties.authUrl + "?appid={appid}&secret={secret}&js_code={jscode}&grant_type=authorization_code",
          WechatAuthInfo.class,
          ImmutableMap.of(
              "appid", wechatProperties.appId,
              "secret", wechatProperties.appSecret,
              "jscode", token));
    } while (info.getErrorCode() == -1);
    if (info.getErrorCode() == 0) {
      return info.getOpenId();
    } else if (info.getErrorCode() == 40029) {
      throw new BadCredentialsException("Wrong wechat mini boardgame token");
    } else if (info.getErrorCode() == 45011) {
      throw new BadCredentialsException("Server busy, try again later");
    } else {
      throw new BadCredentialsException("Unknown error code: " + info.getErrorCode() + ", " + info.getErrorMessage());
    }
  }
}
