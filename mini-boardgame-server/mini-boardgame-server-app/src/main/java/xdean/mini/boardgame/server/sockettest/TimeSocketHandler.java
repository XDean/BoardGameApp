package xdean.mini.boardgame.server.sockettest;

import static xdean.jex.util.lang.ExceptionUtil.uncheck;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.concurrent.TimeUnit;

import org.springframework.stereotype.Service;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketSession;
import org.springframework.web.socket.handler.TextWebSocketHandler;

import io.reactivex.Observable;
import xdean.jex.log.Logable;

@Service
public class TimeSocketHandler extends TextWebSocketHandler implements Logable {

  List<WebSocketSession> sessions = Collections.synchronizedList(new ArrayList<>());

  public TimeSocketHandler() {
    Observable.interval(1, TimeUnit.SECONDS)
        .subscribe(e -> sendTime());
  }

  @Override
  public void afterConnectionEstablished(WebSocketSession session) throws Exception {
    debug("Time socket established: " + session.getId());
    sessions.add(session);
  }

  @Override
  public void afterConnectionClosed(WebSocketSession session, CloseStatus status) throws Exception {
    debug("Time socket closed: " + session.getId() + ". With status:" + status.toString());
    sessions.remove(session);
  }

  @Override
  protected void handleTextMessage(WebSocketSession session, TextMessage message) throws Exception {
    debug("Time socket recive: " + session.getId() + " -> " + message.getPayload());
    sessions.stream()
        .filter(s -> s.isOpen())
        .forEach(s -> uncheck(() -> s.sendMessage(new TextMessage(message.getPayload()))));
  }

  @Override
  public void handleTransportError(WebSocketSession session, Throwable exception) throws Exception {
    debug("Time socket error: " + session.getId(), exception);
  }

  public void sendTime() {
    trace("Send time");
    sessions.stream()
        .filter(s -> s.isOpen())
        .forEach(s -> uncheck(() -> s.sendMessage(new TextMessage("" + System.currentTimeMillis()))));
  }
}
