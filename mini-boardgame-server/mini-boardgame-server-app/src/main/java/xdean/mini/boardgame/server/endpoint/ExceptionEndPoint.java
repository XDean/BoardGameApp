package xdean.mini.boardgame.server.endpoint;

import org.springframework.boot.configurationprocessor.json.JSONObject;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import xdean.mini.boardgame.server.model.exception.MiniBoardgameException;

@RestControllerAdvice
public class ExceptionEndPoint {

  @ExceptionHandler(IllegalArgumentException.class)
  public ResponseEntity<String> handle(IllegalArgumentException e) {
    return handle(MiniBoardgameException.builder().code(HttpStatus.BAD_REQUEST).message(e.getMessage()).build());
  }

  @ExceptionHandler(MiniBoardgameException.class)
  public ResponseEntity<String> handle(MiniBoardgameException e) {
    return ResponseEntity.status(e.getCode()).body(JSONObject.quote(e.getMessage()));
  }
}
