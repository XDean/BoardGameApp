package xdean.mini.boardgame.server.service;

import java.util.Optional;

import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;

public enum UserService {
  ;
  public static Optional<User> getAuthUser() {
    Authentication a = SecurityContextHolder.getContext().getAuthentication();
    if (a != null) {
      Object p = a.getPrincipal();
      if (p instanceof User) {
        return Optional.of((User) p);
      }
    }
    return Optional.empty();
  }
}
