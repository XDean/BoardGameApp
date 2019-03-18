package xdean.mini.boardgame.server.model.exception;

import org.springframework.http.HttpStatus;

import lombok.Builder;
import lombok.EqualsAndHashCode;
import lombok.ToString;
import lombok.Value;

@Builder
@Value
@ToString(callSuper = false)
@EqualsAndHashCode(callSuper = false)
public class MiniBoardgameException extends RuntimeException {
  HttpStatus code;
  String message;
}