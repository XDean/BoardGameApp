package xdean.mini.boardgame.server.endpoint;

import java.time.Duration;
import java.util.concurrent.ThreadLocalRandom;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletResponse;

import org.springframework.http.codec.ServerSentEvent;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.web.bind.annotation.CookieValue;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import io.swagger.annotations.Api;
import reactor.core.publisher.Flux;
import xdean.mini.boardgame.server.security.model.LoginOpenIdResponse;

@RestController
@Api(tags = "Test")
public class TestEndPoint {
  @GetMapping("/hello")
  public String hello(@RequestParam(name = "who", required = false) String who) {
    if (who == null) {
      who = "World";
      Authentication a = SecurityContextHolder.getContext().getAuthentication();
      if (a != null) {
        Object p = a.getPrincipal();
        if (p instanceof User) {
          who = ((User) p).getUsername().toString();
        }
      }
    }
    return "Hello " + who + "!";
  }

  @GetMapping("/hello-cookie")
  public String helloCookie(@CookieValue(name = "who", defaultValue = "World") String who) {
    return "Hello " + who + "!";
  }

  @GetMapping("/hello-cookie-set")
  public String setHelloCookie(HttpServletResponse response, @RequestParam(name = "who", defaultValue = "World") String who) {
    response.addCookie(new Cookie("who", who));
    return "Set to " + who + "!";
  }

  @GetMapping("/flux/ints")
  public Flux<ServerSentEvent<LoginOpenIdResponse>> ints() {
    return Flux.interval(Duration.ofSeconds(1))
        .map(seq -> ServerSentEvent.<LoginOpenIdResponse> builder()
            .event("random")
            .id(seq.toString())
            .data(LoginOpenIdResponse.builder().errorCode(ThreadLocalRandom.current().nextInt()).build())
            .build());
  }
}
