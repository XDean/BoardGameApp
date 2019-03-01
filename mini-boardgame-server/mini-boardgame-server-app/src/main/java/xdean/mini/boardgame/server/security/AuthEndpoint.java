package xdean.mini.boardgame.server.security;

import javax.inject.Inject;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.provisioning.UserDetailsManager;
import org.springframework.security.web.authentication.WebAuthenticationDetails;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import xdean.mini.boardgame.server.security.model.ApplicationUser;
import xdean.mini.boardgame.server.security.model.SignUpResponse;

@RestController
public class AuthEndpoint {

  @Inject
  UserDetailsManager userDetailsManager;

  @Inject
  AuthenticationManager authenticationManager;

  @Inject
  PasswordEncoder passwordEncoder;

  @RequestMapping("/sign-up")
  public SignUpResponse signUp(
      HttpServletRequest request,
      HttpServletResponse response,
      @RequestBody(required = false) ApplicationUser user,
      @RequestParam(required = false) String username,
      @RequestParam(required = false) String password) {
    if (user == null) {
      if (username == null || password == null) {
        return SignUpResponse.builder()
            .success(false)
            .message("Please input username and password")
            .errorCode(SignUpResponse.INPUT_USERNAME_PASSWORD)
            .build();
      } else {
        user = new ApplicationUser(username, password);
      }
    }
    if (!user.getUsername().matches("^(?!_)(?!.*?_$)[a-zA-Z0-9_]+$")) {
      return SignUpResponse.builder()
          .success(false)
          .message("Username should be letter and/or number")
          .errorCode(SignUpResponse.ILLEGAL_USERNAME)
          .build();
    }
    if (!user.getPassword().matches("^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$")) {
      return SignUpResponse.builder()
          .success(false)
          .message("Password should be letter and number")
          .errorCode(SignUpResponse.ILLEGAL_PASSWORD)
          .build();
    }
    boolean exist = userDetailsManager.userExists(user.getUsername());
    if (exist) {
      return SignUpResponse.builder()
          .success(false)
          .message("User name exist")
          .errorCode(SignUpResponse.USERNAME_EXIST)
          .build();
    }
    UserDetails u = User.builder()
        .username(user.getUsername())
        .password(user.getPassword())
        .passwordEncoder(passwordEncoder::encode)
        .authorities("USER")
        .build();
    userDetailsManager.createUser(u);
    authenticateUserAndSetSession(u, user.getPassword(), request);
    return SignUpResponse.builder()
        .success(true)
        .message("Sign up success")
        .build();
  }

  private void authenticateUserAndSetSession(UserDetails user, String rawPassword, HttpServletRequest request) {
    String username = user.getUsername();
    UsernamePasswordAuthenticationToken token = new UsernamePasswordAuthenticationToken(username, rawPassword);

    request.getSession();

    token.setDetails(new WebAuthenticationDetails(request));
    Authentication authenticatedUser = authenticationManager.authenticate(token);

    SecurityContextHolder.getContext().setAuthentication(authenticatedUser);
  }
}