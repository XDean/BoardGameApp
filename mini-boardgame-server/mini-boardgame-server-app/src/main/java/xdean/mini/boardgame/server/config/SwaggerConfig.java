package xdean.mini.boardgame.server.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.MediaType;

import com.google.common.collect.ImmutableSet;

import springfox.documentation.builders.ApiInfoBuilder;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.service.ApiInfo;
import springfox.documentation.service.Contact;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@Configuration
@EnableSwagger2
public class SwaggerConfig {

  @Bean
  public Docket createRestApi() {
    return new Docket(DocumentationType.SWAGGER_2)
        .produces(ImmutableSet.of(MediaType.APPLICATION_JSON.toString()))
        .consumes(ImmutableSet.of(MediaType.APPLICATION_JSON.toString()))
        .apiInfo(apiInfo())
        .select()
        .apis(RequestHandlerSelectors.basePackage("xdean.mini.boardgame"))
        .paths(PathSelectors.any())
        .build();
  }

  private ApiInfo apiInfo() {
    return new ApiInfoBuilder()
        .title("Mini Boardgame API Document")
        .description("https://github.com/XDean/MiniBoardgame")
        .contact(new Contact("XDean", "https://github.com/XDean", "373216024@qq.com"))
        .version("1.0")
        .build();
  }
}