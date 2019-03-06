package xdean.mini.boardgame.server.socket;

import java.util.Map;

import com.fasterxml.jackson.annotation.JsonIgnore;

import lombok.Builder;
import lombok.Builder.Default;
import lombok.Data;

@Data
@Builder(toBuilder = true)
public class WebSocketEvent<T> {

  public static final String ERROR_TOPIC = "ERROR";

  @Default
  @JsonIgnore
  WebSocketSendType type = WebSocketSendType.ALL;

  String topic;
  Map<String, String> attributes;
  T payload;
}
