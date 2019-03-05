package xdean.mini.boardgame.server.socket;

import java.net.URI;
import java.util.AbstractMap.SimpleImmutableEntry;
import java.util.Arrays;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.stream.Collectors;

import javax.inject.Inject;

import org.springframework.stereotype.Component;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketSession;
import org.springframework.web.socket.handler.TextWebSocketHandler;

import com.google.common.base.Strings;

import xdean.jex.extra.tryto.Try;
import xdean.jex.log.Logable;
import xdean.mini.boardgame.server.service.UserService;

@Component
public class GameSocketHandler extends TextWebSocketHandler implements Logable {

  private static final String ID = "id";

  private @Inject UserService userService;

  private final Map<Integer, GameRoomSocketHandler> rooms = new ConcurrentHashMap<>();

  @Override
  public void afterConnectionEstablished(WebSocketSession session) throws Exception {
    Map<String, List<String>> params = splitQuery(session.getUri());
    List<String> values = params.get(ID);
    if (values.size() != 1) {
      session.close(CloseStatus.BAD_DATA.withReason("Must have 'id' param"));
    }
    String idStr = values.get(0);
    Try.to(() -> Integer.parseInt(idStr)).toOptional().ifPresent(id -> {
      trace(session.getRemoteAddress() + " connect to room " + id);
      session.getAttributes().put(ID, id);
      GameRoomSocketHandler room = rooms.computeIfAbsent(id, GameRoomSocketHandler::new);
      room.sessions.add(session);
    });
  }

  @Override
  public void afterConnectionClosed(WebSocketSession session, CloseStatus status) throws Exception {
    Integer id = (Integer) session.getAttributes().get(ID);
    if (id != null) {
      GameRoomSocketHandler room = rooms.get(id);
      if (room != null) {
        room.sessions.remove(session);
      }
    }
  }

  @Override
  public void handleTransportError(WebSocketSession session, Throwable exception) throws Exception {
    super.handleTransportError(session, exception);
  }

  @Override
  protected void handleTextMessage(WebSocketSession session, TextMessage message) throws Exception {
    super.handleTextMessage(session, message);
  }

  /**
   * @author Pr0gr4mm3r@stackoverflow
   */
  private static Map<String, List<String>> splitQuery(URI uri) {
    if (Strings.isNullOrEmpty(uri.getQuery())) {
      return Collections.emptyMap();
    }
    return Arrays.stream(uri.getQuery().split("&"))
        .map(e -> splitQueryParameter(e))
        .collect(Collectors.groupingBy(SimpleImmutableEntry::getKey, LinkedHashMap::new,
            Collectors.mapping(Map.Entry::getValue, Collectors.toList())));
  }

  private static SimpleImmutableEntry<String, String> splitQueryParameter(String it) {
    final int idx = it.indexOf("=");
    final String key = idx > 0 ? it.substring(0, idx) : it;
    final String value = idx > 0 && it.length() > idx + 1 ? it.substring(idx + 1) : null;
    return new SimpleImmutableEntry<>(key, value);
  }
}
