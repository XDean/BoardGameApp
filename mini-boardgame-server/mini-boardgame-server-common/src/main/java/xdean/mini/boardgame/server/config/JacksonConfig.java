package xdean.mini.boardgame.server.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.fasterxml.jackson.databind.ObjectMapper;

@Configuration
public class JacksonConfig {

  public static final ObjectMapper OBJECT_MAPPER = new ObjectMapper();

  @Bean
  public ObjectMapper objectMapper() {
    return OBJECT_MAPPER;
  }
}
