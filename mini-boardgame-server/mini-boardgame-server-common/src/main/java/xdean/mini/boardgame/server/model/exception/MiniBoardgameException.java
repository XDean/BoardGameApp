package xdean.mini.boardgame.server.model.exception;

import lombok.Builder;

public class MiniBoardgameException extends RuntimeException {

  private final MiniBoardgameErrorCode code;

  @Builder
  public MiniBoardgameException(String message, Throwable cause, MiniBoardgameErrorCode code) {
    super(message, cause);
    this.code = code;
  }

  public MiniBoardgameErrorCode getCode() {
    return code;
  }

  public static MiniBoardgameExceptionBuilder internal() {
    return builder().code(StandardErrorCode.INTERNAL);
  }

  public static MiniBoardgameExceptionBuilder unkown() {
    return builder().code(StandardErrorCode.UNKNOWN);
  }

  public static class MiniBoardgameExceptionBuilder {
    public MiniBoardgameExceptionBuilder cause(Throwable cause) {
      this.cause = cause;
      if (this.message == null) {
        this.message = cause.getMessage();
      }
      if ((this.code == null || this.code == StandardErrorCode.UNKNOWN) && cause instanceof MiniBoardgameException) {
        this.code = ((MiniBoardgameException) cause).getCode();
      }
      return this;
    }
  }
}