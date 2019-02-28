package xdean.wechat.mini.boardgame.server.security.config;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;

import xdean.wechat.mini.boardgame.server.security.SecurityProperties;

@EnableWebSecurity
@EnableConfigurationProperties
public class SecurityConfig extends WebSecurityConfigurerAdapter {
  @Override
  protected void configure(HttpSecurity http) throws Exception {
    http
        .csrf().disable()
        .authorizeRequests()
        .antMatchers("/sign-up", "/login").permitAll()
        .anyRequest().authenticated()
        .and()
        .addFilter(new JWTAuthorizationFilter(authenticationManager(), securityProperties()));
//        // this disables session creation on Spring Security
//        .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS);
  }

  @Bean
  public BCryptPasswordEncoder bCryptPasswordEncoder() {
    return new BCryptPasswordEncoder();
  }

  @Bean
  @ConfigurationProperties(prefix = "mini-boardgame.security")
  public SecurityProperties securityProperties() {
    return new SecurityProperties();
  }
}