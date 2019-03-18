package xdean.mini.boardgame.server.model.exception;

import org.springframework.http.HttpStatus;

import lombok.Builder;

@Builder
public class MiniBoardgameException extends RuntimeException {
  HttpStatus code;
  String message;
}