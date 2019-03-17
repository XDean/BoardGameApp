package xdean.mini.boardgame.server.socket;

import java.util.Map;

import com.fasterxml.jackson.annotation.JsonIgnore;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Builder.Default;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.Singular;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder(toBuilder = true)
public class WebSocketEvent<T> {

  @Default
  @JsonIgnore
  WebSocketSendType type = WebSocketSendType.ALL;

  String topic;

  @Singular
  Map<String, Object> attributes;

  T payload;
}
