package xdean.mini.boardgame.server.endpoint;

import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import xdean.mini.boardgame.server.model.param.IllegalArgumentResponse;

@RestControllerAdvice
public class ExceptionEndPoint {

  @ExceptionHandler(IllegalArgumentException.class)
  public IllegalArgumentResponse handle(IllegalArgumentException e) {
    return new IllegalArgumentResponse(e.getMessage());
  }
}
