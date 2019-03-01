package xdean.mini.boardgame.server.security.handler;

import java.io.IOException;
import java.nio.file.ProviderNotFoundException;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.authentication.AbstractAuthenticationProcessingFilter;
import org.springframework.security.web.util.matcher.AntPathRequestMatcher;

import com.google.common.base.Strings;

import xdean.jex.log.Logable;

public class OpenIdAuthFilter extends AbstractAuthenticationProcessingFilter implements Logable {

  @Autowired(required = false)
  List<OpenIdAuthProvider> providers = Collections.emptyList();

  public OpenIdAuthFilter() {
    super(new AntPathRequestMatcher("/login", "POST"));
  }

  @Override
  public Authentication attemptAuthentication(HttpServletRequest request, HttpServletResponse response)
      throws AuthenticationException, IOException, ServletException {
    String type = request.getParameter("type");
    if ("openid".equals(type)) {
      return null;
    }
    String provider = Strings.nullToEmpty(request.getParameter("provider"));
    String token = Strings.nullToEmpty(request.getParameter("token"));

    List<OpenIdAuthProvider> findProviders = providers.stream().filter(p -> p.name().equals(provider))
        .collect(Collectors.toList());
    if (findProviders.isEmpty()) {
      throw new ProviderNotFoundException("There is no provider support: " + provider);
    }
    for (int i = 0; i < findProviders.size(); i++) {
      OpenIdAuthProvider p = findProviders.get(i);
      try {
        String result = p.attemptAuthentication(token);
        if (result != null) {
          return new UsernamePasswordAuthenticationToken(result, null);
        }
      } catch (AuthenticationException e) {
        trace().log("Fail to authenticate: " + token, e);
      }
    }
    throw new BadCredentialsException("Bad Credentials");
  }
}
