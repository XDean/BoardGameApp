package xdean.mini.boardgame.server.security;

import javax.inject.Inject;
import javax.sql.DataSource;

import org.springframework.boot.autoconfigure.jdbc.DataSourceProperties;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.provisioning.UserDetailsManager;
import org.springframework.security.web.access.intercept.FilterSecurityInterceptor;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

import xdean.mini.boardgame.server.handler.DispatchLoginHandler;
import xdean.mini.boardgame.server.handler.DispatchLogoutHandler;
import xdean.mini.boardgame.server.security.handler.AuthenticationInfoHandler;
import xdean.mini.boardgame.server.security.handler.TokenAuthenticationProvider;
import xdean.mini.boardgame.server.security.handler.UserDetailsManagerImpl;
import xdean.mini.boardgame.server.security.model.SecurityProperties;

@Configuration
@EnableWebSecurity
@EnableConfigurationProperties
@EnableGlobalMethodSecurity(securedEnabled = true)
public class SecurityConfig extends WebSecurityConfigurerAdapter {
  @Inject
  DataSource dataSource;

  @Inject
  DataSourceProperties dataSourceProperties;

  @Inject
  DispatchLoginHandler loginHandler;

  @Inject
  DispatchLogoutHandler logoutHandler;

  @Inject
  DispatchAuthenticationHandler authenticationHandler;

  @Inject
  TokenAuthenticationProvider tokenAuthProvider;

  @Inject
  AuthenticationInfoHandler authInfoHandler;

  @Override
  protected void configure(HttpSecurity http) throws Exception {
    loginHandler.setDefaultTargetUrl("/hello");
    http
        .csrf().disable()
        .addFilterAfter(authInfoHandler, FilterSecurityInterceptor.class)
        .addFilterBefore(tokenAuthProvider, UsernamePasswordAuthenticationFilter.class)
        .authorizeRequests()
        .antMatchers("/sign-up", "/login**").permitAll()
        .antMatchers("/**/favicon.ico", "/webjars/**").permitAll()
        .anyRequest().authenticated()
        .and()
        .formLogin()
        .successHandler(loginHandler)
        .and()
        .logout()
        .addLogoutHandler(logoutHandler)
        .deleteCookies(TokenAuthenticationProvider.ACCESS_TOKEN)
        .and();
  }

  @Override
  protected void configure(AuthenticationManagerBuilder auth) throws Exception {
    auth
        .authenticationEventPublisher(authenticationHandler)
        .authenticationProvider(tokenAuthProvider)
        .userDetailsService(userDetailsManager())
        .passwordEncoder(passwordEncoder());
    if (!userDetailsManager().userExists("admin")) {
      userDetailsManager().createUser(User.builder()
          .username("admin")
          .password(dataSourceProperties.getPassword())
          .passwordEncoder(passwordEncoder()::encode)
          .authorities("ROLE_ADMIN", "ROLE_USER")
          .build());
    }
  }

  @Bean
  @Override
  protected AuthenticationManager authenticationManager() throws Exception {
    return super.authenticationManager();
  }

  @Bean
  public UserDetailsManager userDetailsManager() {
    return new UserDetailsManagerImpl();
  }

  @Bean
  public PasswordEncoder passwordEncoder() {
    return new BCryptPasswordEncoder();
  }

  @Bean
  @ConfigurationProperties(prefix = "mini-boardgame.security")
  public SecurityProperties securityProperties() {
    return new SecurityProperties();
  }
}