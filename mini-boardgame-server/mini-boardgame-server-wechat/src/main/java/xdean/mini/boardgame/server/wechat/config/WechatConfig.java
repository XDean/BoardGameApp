package xdean.mini.boardgame.server.wechat.config;

import java.util.Arrays;

import javax.inject.Inject;

import org.springframework.context.annotation.Configuration;
import org.springframework.http.MediaType;
import org.springframework.http.converter.json.MappingJackson2HttpMessageConverter;
import org.springframework.web.client.RestTemplate;

@Configuration
public class WechatConfig {

  @Inject
  public void initWechatConverter(RestTemplate restTemplate) {
    MappingJackson2HttpMessageConverter c = new MappingJackson2HttpMessageConverter();
    c.setSupportedMediaTypes(Arrays.asList(MediaType.TEXT_PLAIN));
    restTemplate.getMessageConverters().add(c);
  }
}
