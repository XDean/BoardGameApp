package xdean.mini.boardgame.server.security.handler;

import java.util.Optional;
import java.util.stream.Collectors;

import javax.inject.Inject;

import org.springframework.security.access.AccessDeniedException;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.provisioning.UserDetailsManager;
import org.springframework.stereotype.Component;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.service.UserDataService;

@Component
public class UserDetailsManagerImpl implements UserDetailsManager {
  @Inject
  UserDataService userDataService;

  @Inject
  AuthenticationManager authenticationManager;

  @Override
  public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
    Optional<UserEntity> user = userDataService.findUserByUsername(username);
    if (user.isPresent()) {
      UserEntity u = user.get();
      return User.builder()
          .username(u.getUsername())
          .password(u.getPassword())
          .authorities(u.getAuthorities().stream().toArray(String[]::new))
          .disabled(!u.isEnabled())
          .build();
    } else {
      throw new UsernameNotFoundException("There is no user named: " + username);
    }
  }

  @Override
  public void createUser(UserDetails user) {
    if (userDataService.userExist(user.getUsername())) {
      throw new IllegalArgumentException("Username exist: " + user.getUsername());
    }
    updateUser(user);
  }

  @Override
  public void updateUser(UserDetails user) {
    if (!userDataService.userExist(user.getUsername())) {
      throw new IllegalArgumentException("Username doesn't exist: " + user.getUsername());
    }
    saveUser(user);
  }

  @Override
  public void deleteUser(String username) {
    userDataService.delete(username);
  }

  @Override
  public void changePassword(String oldPassword, String newPassword) {
    Authentication currentUser = SecurityContextHolder.getContext()
        .getAuthentication();
    if (currentUser == null) {
      throw new AccessDeniedException(
          "Can't change password as no Authentication object found in context for current user.");
    }
    String username = currentUser.getName();
    authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(
        username, oldPassword));
    userDataService.changePassword(username, newPassword);
    SecurityContextHolder.getContext().setAuthentication(
        createNewAuthentication(currentUser, newPassword));
  }

  @Override
  public boolean userExists(String username) {
    return userDataService.userExist(username);
  }

  private void saveUser(UserDetails user) {
    UserEntity u = UserEntity.builder()
        .username(user.getUsername())
        .password(user.getPassword())
        .enabled(user.isEnabled())
        .authorities(user.getAuthorities().stream().map(a -> a.getAuthority()).collect(Collectors.toList()))
        .build();
    userDataService.save(u);
  }

  private Authentication createNewAuthentication(Authentication currentAuth,
      String newPassword) {
    UserDetails user = loadUserByUsername(currentAuth.getName());

    UsernamePasswordAuthenticationToken newAuthentication = new UsernamePasswordAuthenticationToken(
        user, null, user.getAuthorities());
    newAuthentication.setDetails(currentAuth.getDetails());

    return newAuthentication;
  }
}
