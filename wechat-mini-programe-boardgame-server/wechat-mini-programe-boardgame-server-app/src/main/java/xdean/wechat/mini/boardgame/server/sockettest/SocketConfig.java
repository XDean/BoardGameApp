package xdean.wechat.mini.boardgame.server.sockettest;

import javax.inject.Inject;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.socket.config.annotation.EnableWebSocket;
import org.springframework.web.socket.config.annotation.WebSocketConfigurer;
import org.springframework.web.socket.config.annotation.WebSocketHandlerRegistry;

import xdean.wechat.mini.boardgame.server.controller.TimeSocketHandler;

@Configuration
@EnableWebSocket
public class SocketConfig implements WebSocketConfigurer {
  @Inject
  TimeSocketHandler handler;

  @Override
  public void registerWebSocketHandlers(WebSocketHandlerRegistry registry) {
    registry
        .addHandler(handler, "/socket-test")
        .setAllowedOrigins("*");
  }
}