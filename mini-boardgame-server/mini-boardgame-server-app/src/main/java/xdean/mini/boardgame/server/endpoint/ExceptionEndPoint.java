package xdean.mini.boardgame.server.endpoint;

import org.springframework.boot.configurationprocessor.json.JSONObject;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import xdean.mini.boardgame.server.model.exception.MiniBoardgameException;
import xdean.mini.boardgame.server.model.param.IllegalArgumentResponse;

@RestControllerAdvice
public class ExceptionEndPoint {

  @ExceptionHandler(IllegalArgumentException.class)
  public IllegalArgumentResponse handle(IllegalArgumentException e) {
    return new IllegalArgumentResponse(e.getMessage());
  }

  @ExceptionHandler(MiniBoardgameException.class)
  public ResponseEntity<?> handle(MiniBoardgameException e) {
    return ResponseEntity.status(e.getCode()).body(JSONObject.quote(e.getMessage()));
  }
}
