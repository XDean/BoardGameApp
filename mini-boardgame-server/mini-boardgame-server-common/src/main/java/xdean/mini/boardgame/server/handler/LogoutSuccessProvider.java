package xdean.mini.boardgame.server.handler;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public interface LogoutSuccessProvider {
  void afterSuccess(HttpServletRequest request, HttpServletResponse response, String username);
}
