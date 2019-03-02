package xdean.mini.boardgame.server.endpoint;

import java.util.concurrent.TimeUnit;

import javax.inject.Inject;

import org.springframework.boot.SpringApplication;
import org.springframework.context.ApplicationContext;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import io.reactivex.schedulers.Schedulers;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import io.swagger.annotations.Authorization;
import xdean.mini.boardgame.server.security.annotation.AdminAuth;

@RestController
@Api(tags = "Admin", authorizations = @Authorization(value = "ADMIN"))
public class SystemEndPoint {

  private @Inject ApplicationContext applicationContext;

  @AdminAuth
  @GetMapping("/shutdown")
  @ApiOperation("Shutdown the server")
  public String shutdown(@RequestParam(name = "delay", required = false, defaultValue = "1000") int delay) {
    Schedulers.io().scheduleDirect(() -> SpringApplication.exit(applicationContext), delay, TimeUnit.MILLISECONDS);
    return "SHUTDOWN";
  }
}
