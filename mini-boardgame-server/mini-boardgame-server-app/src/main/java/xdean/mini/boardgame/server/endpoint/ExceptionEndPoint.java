package xdean.mini.boardgame.server.endpoint;

import java.util.HashMap;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import xdean.mini.boardgame.server.model.exception.MiniBoardgameException;

@RestControllerAdvice
public class ExceptionEndPoint {

  @ExceptionHandler(IllegalArgumentException.class)
  public ResponseEntity<?> handle(IllegalArgumentException e) {
    return handle(MiniBoardgameException.builder().code(HttpStatus.BAD_REQUEST).message(e.getMessage()).build());
  }

  @ExceptionHandler(MiniBoardgameException.class)
  public ResponseEntity<?> handle(MiniBoardgameException e) {
    HashMap<Object, Object> map = new HashMap<>();
    map.put("code", e.getCode().value());
    map.put("message", e.getMessage());
    map.put("key", e.getErrorKey());
    map.put("details", e.getDetails());
    return ResponseEntity.status(e.getCode()).body(map);
  }
}
