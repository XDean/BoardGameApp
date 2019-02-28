package xdean.wechat.mini.boardgame.server.security;

import static com.auth0.jwt.algorithms.Algorithm.HMAC512;

import java.util.Date;

import javax.inject.Inject;
import javax.servlet.http.HttpServletResponse;

import org.springframework.http.HttpHeaders;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.auth0.jwt.JWT;

import xdean.wechat.mini.boardgame.server.security.model.ApplicationUser;
import xdean.wechat.mini.boardgame.server.security.model.LoginResponse;
import xdean.wechat.mini.boardgame.server.security.model.SignUpResponse;
import xdean.wechat.mini.boardgame.server.security.service.ApplicationUserRepository;

@RestController
public class AuthEndpoint {

  @Inject
  ApplicationUserRepository applicationUserRepository;

  @Inject
  BCryptPasswordEncoder bCryptPasswordEncoder;

  @Inject
  SecurityProperties properties;

  @RequestMapping("/sign-up")
  public SignUpResponse signUp(
      HttpServletResponse res,
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
        user = new ApplicationUser(0, username, password);
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
    ApplicationUser find = applicationUserRepository.findByUsername(user.getUsername());
    if (find != null) {
      return SignUpResponse.builder()
          .success(false)
          .message("User name exist")
          .errorCode(SignUpResponse.USERNAME_EXIST)
          .build();
    }
    user.setPassword(bCryptPasswordEncoder.encode(user.getPassword()));
    applicationUserRepository.save(user);
    addToken(res, user);
    return SignUpResponse.builder()
        .success(true)
        .message("Sign up success")
        .build();
  }

  @RequestMapping("/login")
  public LoginResponse login(
      HttpServletResponse res,
      @RequestBody(required = false) ApplicationUser user,
      @RequestParam(required = false) String username,
      @RequestParam(required = false) String password) {
    if (user == null) {
      if (username == null || password == null) {
        return LoginResponse.builder()
            .success(false)
            .message("Please input username and password or goto sign-up")
            .errorCode(LoginResponse.INPUT_USERNAME_PASSWORD)
            .build();
      } else {
        user = new ApplicationUser(0, username, password);
      }
    }
    ApplicationUser find = applicationUserRepository.findByUsername(user.getUsername());
    if (find != null) {
      if (bCryptPasswordEncoder.matches(user.getPassword(), find.getPassword())) {
        addToken(res, user);
        return LoginResponse.builder()
            .success(true)
            .message("Login success")
            .build();
      }
    }
    return LoginResponse.builder()
        .success(false)
        .message("Wrong password or user not exist")
        .errorCode(LoginResponse.WRONG_PASSWORD_OR_USERNAME_NOT_EXIST)
        .build();
  }

  private void addToken(HttpServletResponse res, ApplicationUser user) {
    String token = JWT.create()
        .withSubject(user.getUsername())
        .withExpiresAt(new Date(System.currentTimeMillis() + properties.getExpirationTime()))
        .sign(HMAC512(properties.getSecretKey().getBytes()));
    res.addHeader(HttpHeaders.AUTHORIZATION, properties.getTokenPrefix() + token);
  }
}