package xdean.mini.boardgame.server.service.impl;

import java.util.Optional;

import javax.inject.Inject;

import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.stereotype.Service;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.service.UserEntityRepo;
import xdean.mini.boardgame.server.service.UserService;

@Service
public class UserServiceImpl implements UserService {

  @Inject
  UserEntityRepo userEntityRepo;

  @Override
  public Optional<UserEntity> getCurrentUser() {
    Authentication a = SecurityContextHolder.getContext().getAuthentication();
    if (a != null) {
      Object p = a.getPrincipal();
      if (p instanceof User) {
        String username = ((User) p).getUsername();
        return getUserByUsername(username);
      }
    }
    return Optional.empty();
  }

  @Override
  public Optional<UserEntity> getUserByUsername(String username) {
    return userEntityRepo.findByUsername(username);
  }
}
