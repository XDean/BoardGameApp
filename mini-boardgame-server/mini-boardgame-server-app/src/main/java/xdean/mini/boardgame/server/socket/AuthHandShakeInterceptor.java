package xdean.mini.boardgame.server.socket;

import java.util.List;
import java.util.Map;

import org.springframework.http.HttpHeaders;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.server.HandshakeInterceptor;

import xdean.mini.boardgame.server.model.CommonConstants.AttrKey;

@Component
public class AuthHandShakeInterceptor implements HandshakeInterceptor {

  @Override
  public boolean beforeHandshake(ServerHttpRequest request, ServerHttpResponse response, WebSocketHandler wsHandler,
      Map<String, Object> attributes) throws Exception {
    List<String> token = request.getHeaders().get(HttpHeaders.AUTHORIZATION);
    if (!token.isEmpty()) {
      attributes.put(AttrKey.ACCESS_TOKEN, token.get(0));
    }
    return true;
  }

  @Override
  public void afterHandshake(ServerHttpRequest request, ServerHttpResponse response, WebSocketHandler wsHandler,
      Exception exception) {
  }
}
