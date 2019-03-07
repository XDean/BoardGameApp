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
import org.springframework.security.provisioning.JdbcUserDetailsManager;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

import xdean.mini.boardgame.server.handler.DispatchLoginHandler;
import xdean.mini.boardgame.server.handler.DispatchLogoutHandler;
import xdean.mini.boardgame.server.security.handler.JwtTokenHandler;
import xdean.mini.boardgame.server.security.model.SecurityProperties;

@Configuration
@EnableWebSecurity
@EnableConfigurationProperties
@EnableGlobalMethodSecurity(prePostEnabled = true, securedEnabled = true, jsr250Enabled = true)
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
  JwtTokenHandler jwtTokenHandler;

  @Override
  protected void configure(HttpSecurity http) throws Exception {
    loginHandler.setDefaultTargetUrl("/hello");
    logoutHandler.setDefaultTargetUrl("/login?logout");
    http
        .csrf().disable()
        .addFilterBefore(jwtTokenHandler, UsernamePasswordAuthenticationFilter.class)
        .authorizeRequests()
        .antMatchers("/sign-up", "/login**").permitAll()
        .antMatchers("/**/favicon.ico", "/webjars/**").permitAll()
        .anyRequest().authenticated()
        .and()
        .formLogin()
        .successHandler(loginHandler)
        .and()
        .logout()
        .logoutSuccessHandler(logoutHandler)
        .deleteCookies(JwtTokenHandler.JWT_TOKEN);
  }

  @Override
  protected void configure(AuthenticationManagerBuilder auth) throws Exception {
    JdbcUserDetailsManager m = userDetailsManager();
    auth
        .authenticationProvider(jwtTokenHandler)
        .userDetailsService(m)
        .passwordEncoder(passwordEncoder());
    if (!m.userExists("admin")) {
      m.createUser(User.builder()
          .username("admin")
          .password(dataSourceProperties.getPassword())
          .passwordEncoder(passwordEncoder()::encode)
          .authorities("USER", "ADMIN")
          .build());
    }
  }

  @Bean
  @Override
  protected AuthenticationManager authenticationManager() throws Exception {
    return super.authenticationManager();
  }

  @Bean
  public JdbcUserDetailsManager userDetailsManager() {
    JdbcUserDetailsManager manager = new JdbcUserDetailsManager();
    manager.setDataSource(dataSource);
    return manager;
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