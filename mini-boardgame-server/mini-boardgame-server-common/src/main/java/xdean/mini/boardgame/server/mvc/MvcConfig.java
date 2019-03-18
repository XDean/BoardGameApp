package xdean.mini.boardgame.server.mvc;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.ViewResolver;
import org.springframework.web.servlet.config.annotation.DefaultServletHandlerConfigurer;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;
import org.springframework.web.servlet.view.AbstractTemplateViewResolver;

@Configuration
public class MvcConfig implements WebMvcConfigurer {
  @Override
  public void configureDefaultServletHandling(DefaultServletHandlerConfigurer configurer) {
    configurer.enable();
  }

  @Bean
  public ViewResolver markdownResolver() {
    return new AbstractTemplateViewResolver() {
      {
        setViewClass(requiredViewClass());
        setViewNames("*.md");
      }

      @Override
      protected Class<?> requiredViewClass() {
        return MarkdownView.class;
      }
    };
  }
}
