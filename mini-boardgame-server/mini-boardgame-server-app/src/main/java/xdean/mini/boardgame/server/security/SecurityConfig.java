package xdean.mini.boardgame.server.security;

import javax.inject.Inject;
import javax.sql.DataSource;

import org.springframework.boot.autoconfigure.jdbc.DataSourceProperties;
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

@Configuration
@EnableWebSecurity
@EnableGlobalMethodSecurity(prePostEnabled = true, securedEnabled = true, jsr250Enabled = true)
public class SecurityConfig extends WebSecurityConfigurerAdapter {
  @Inject
  DataSource dataSource;

  @Inject
  DataSourceProperties dataSourceProperties;

  @Override
  protected void configure(HttpSecurity http) throws Exception {
    http
        .csrf().disable()
        .authorizeRequests()
        .antMatchers("/sign-up", "/login**").permitAll()
        .antMatchers("/**/favicon.ico", "/webjars/**").permitAll()
        .anyRequest().authenticated()
        .and()
        .formLogin().defaultSuccessUrl("/hello")
        .and().logout();
  }

  @Override
  protected void configure(AuthenticationManagerBuilder auth) throws Exception {
    JdbcUserDetailsManager m = userDetailsManager();
    auth
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
}