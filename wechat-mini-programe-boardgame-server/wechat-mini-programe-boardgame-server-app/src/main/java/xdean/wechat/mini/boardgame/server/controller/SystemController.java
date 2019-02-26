package xdean.wechat.mini.boardgame.server.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SystemController {

  @GetMapping("/hello")
  public String hello(@RequestParam(name = "who", defaultValue = "World") String who) {
    return "Hello " + who + "!";
  }
}
