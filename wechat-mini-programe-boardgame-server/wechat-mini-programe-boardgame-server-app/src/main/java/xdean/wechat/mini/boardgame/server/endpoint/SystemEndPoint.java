package xdean.wechat.mini.boardgame.server.endpoint;

import java.util.concurrent.TimeUnit;

import javax.inject.Inject;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletResponse;

import org.springframework.boot.SpringApplication;
import org.springframework.context.ApplicationContext;
import org.springframework.web.bind.annotation.CookieValue;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import io.reactivex.schedulers.Schedulers;

@RestController
public class SystemEndPoint {

  private @Inject ApplicationContext applicationContext;

  @GetMapping("/hello")
  public String hello(@RequestParam(name = "who", defaultValue = "World") String who) {
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

  @GetMapping("/shutdown")
  public String shutdown(@RequestParam(name = "delay", required = false, defaultValue = "1000") int delay) {
    Schedulers.io().scheduleDirect(() -> SpringApplication.exit(applicationContext), delay, TimeUnit.MILLISECONDS);
    return "SHUTDOWN";
  }
}