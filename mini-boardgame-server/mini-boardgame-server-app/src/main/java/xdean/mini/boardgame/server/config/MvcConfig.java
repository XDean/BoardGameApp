package xdean.mini.boardgame.server.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.ViewResolver;
import org.springframework.web.servlet.config.annotation.DefaultServletHandlerConfigurer;
import org.springframework.web.servlet.config.annotation.ViewControllerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;
import org.springframework.web.servlet.view.AbstractTemplateViewResolver;

import xdean.mini.boardgame.server.mvc.MarkdownView;

@Configuration
public class MvcConfig implements WebMvcConfigurer {
  @Override
  public void configureDefaultServletHandling(DefaultServletHandlerConfigurer configurer) {
    configurer.enable();
  }

  @Override
  public void addViewControllers(ViewControllerRegistry registry) {
    registry.addViewController("/doc/**");
  }

  @Bean
  public ViewResolver markdownResolver() {
    return new AbstractTemplateViewResolver() {
      {
        setViewClass(requiredViewClass());
        setSuffix(".md");
      }

      @Override
      protected Class<?> requiredViewClass() {
        return MarkdownView.class;
      }
    };
  }
}
