package xdean.mini.boardgame.server.handler;

import java.util.Collections;
import java.util.List;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.web.authentication.logout.LogoutHandler;
import org.springframework.stereotype.Component;

@Component
public class DispatchLogoutHandler implements LogoutHandler {

  @Autowired(required = false)
  List<LogoutSuccessProvider> handlers = Collections.emptyList();

  public void afterSuccess(HttpServletRequest request, HttpServletResponse response) {
    handlers.forEach(h -> h.afterSuccessLogout(request, response));
  }

  @Override
  public void logout(HttpServletRequest request, HttpServletResponse response, Authentication authentication) {
    afterSuccess(request, response);
  }
}
