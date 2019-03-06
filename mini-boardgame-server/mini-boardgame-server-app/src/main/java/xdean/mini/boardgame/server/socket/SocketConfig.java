package xdean.mini.boardgame.server.socket;

import static xdean.jex.util.lang.ExceptionUtil.uncheck;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

import javax.inject.Inject;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.socket.config.annotation.EnableWebSocket;
import org.springframework.web.socket.config.annotation.WebSocketConfigurer;
import org.springframework.web.socket.config.annotation.WebSocketHandlerRegistry;
import org.springframework.web.socket.server.support.HttpSessionHandshakeInterceptor;

import xdean.mini.boardgame.server.annotation.Tag;
import xdean.mini.boardgame.server.model.GameConstants;

@Configuration
@EnableWebSocket
public class SocketConfig implements WebSocketConfigurer, GameConstants {
  @Inject
  TimeSocketHandler handler;

  @Inject
  GameSocketHandler gameHandler;

  @Override
  public void registerWebSocketHandlers(WebSocketHandlerRegistry registry) {
    List<String> attrs = Arrays.stream(GameConstants.AttrKey.class.getFields())
        .filter(f -> {
          Tag tag = f.getAnnotation(Tag.class);
          return f.getType() == String.class && tag != null && Arrays.asList(tag.value()).contains(TagKey.SOCKET_ATTR);
        })
        .map(f -> uncheck(() -> (String) f.get(null)))
        .collect(Collectors.toList());
    registry
        .addHandler(handler, "/socket-test/**")
        .addHandler(gameHandler, "/game/room")
        .addInterceptors(new HttpSessionHandshakeInterceptor(attrs))
        .setAllowedOrigins("*");
  }
}