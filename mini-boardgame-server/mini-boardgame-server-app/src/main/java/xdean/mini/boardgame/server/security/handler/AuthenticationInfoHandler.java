package xdean.mini.boardgame.server.security.handler;

import java.io.IOException;
import java.util.Optional;

import javax.inject.Inject;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import xdean.mini.boardgame.server.handler.LoginSuccessProvider;
import xdean.mini.boardgame.server.model.CommonConstants.AttrKey;
import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.service.GameDataService;
import xdean.mini.boardgame.server.service.UserDataService;

@Component
public class AuthenticationInfoHandler extends OncePerRequestFilter implements LoginSuccessProvider {

  private @Inject UserDataService userService;
  private @Inject GameDataService gameService;

  @Override
  protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain)
      throws ServletException, IOException {
    if (request.getSession().getAttribute(AttrKey.USER_ID) == null) {
      addAuthInfos(request);
    }
    filterChain.doFilter(request, response);
  }

  @Override
  public void afterSuccessLogin(HttpServletRequest request, HttpServletResponse response, String username) {
    addAuthInfos(request);
  }

  private void addAuthInfos(HttpServletRequest request) {
    Optional<UserEntity> user = userService.getCurrentUser();
    if (user.isPresent()) {
      int id = user.get().getId();
      request.getSession().setAttribute(AttrKey.USER_ID, id);
      gameService.findPlayer(id).getRoom().ifPresent(e -> request.getSession().setAttribute(AttrKey.ROOM, e));
    }
  }
}
