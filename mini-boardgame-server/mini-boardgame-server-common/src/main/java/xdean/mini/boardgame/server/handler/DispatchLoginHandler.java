package xdean.mini.boardgame.server.handler;

import java.io.IOException;
import java.util.Collections;
import java.util.List;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.web.authentication.AuthenticationSuccessHandler;
import org.springframework.security.web.authentication.SavedRequestAwareAuthenticationSuccessHandler;
import org.springframework.stereotype.Component;

@Component
public class DispatchLoginHandler extends SavedRequestAwareAuthenticationSuccessHandler implements AuthenticationSuccessHandler {

  @Autowired(required = false)
  List<LoginSuccessProvider> handlers = Collections.emptyList();

  public void afterSuccess(HttpServletRequest request, HttpServletResponse response, String username) {
    handlers.forEach(h -> h.afterSuccessLogin(request, response, username));
  }

  @Override
  public void onAuthenticationSuccess(HttpServletRequest request, HttpServletResponse response,
      Authentication authentication) throws IOException, ServletException {
    Object principal = authentication.getPrincipal();
    if (principal instanceof UserDetails) {
      afterSuccess(request, response, ((UserDetails) principal).getUsername());
    }
    super.onAuthenticationSuccess(request, response, authentication);
  }
}
