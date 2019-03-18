package xdean.mini.boardgame.server.mvc;

import org.apache.catalina.servlet4preview.http.HttpServletRequest;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.servlet.HandlerMapping;

@Controller
public class MvcController {

  @RequestMapping({ "/**/*.md", "/**/*.json" })
  public String mvc(HttpServletRequest request) {
    return request.getAttribute(HandlerMapping.PATH_WITHIN_HANDLER_MAPPING_ATTRIBUTE).toString();
  }
}
