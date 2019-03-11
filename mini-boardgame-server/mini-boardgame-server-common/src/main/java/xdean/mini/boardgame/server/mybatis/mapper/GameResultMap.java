package xdean.mini.boardgame.server.mybatis.mapper;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import xdean.mybatis.extension.ConfigurationInitializer;

@Configuration
public class GameResultMap {

  @Bean
  public ConfigurationInitializer gameMapper() {
    // return InitResultMap.create(type)
    return null;
  }
}
