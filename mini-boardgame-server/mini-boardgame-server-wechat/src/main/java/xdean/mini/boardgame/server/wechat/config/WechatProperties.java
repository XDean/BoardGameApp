package xdean.mini.boardgame.server.wechat.config;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.PropertySource;
import org.springframework.stereotype.Component;

import lombok.Data;
import xdean.mini.boardgame.server.handler.YamlPropertySourceFactory;

@Data
@Component
@PropertySource(value = "wechat-config.yml", factory = YamlPropertySourceFactory.class)
@ConfigurationProperties(prefix = "mini-boardgame.wechat")
public class WechatProperties {
  public String appId;
  public String appSecret;
  public String authUrl = "https://api.weixin.qq.com/sns/jscode2session";
}
