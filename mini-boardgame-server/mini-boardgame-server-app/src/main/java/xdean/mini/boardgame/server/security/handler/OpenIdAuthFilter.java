//package xdean.mini.boardgame.server.security.handler;
//
//import java.io.IOException;
//import java.nio.file.ProviderNotFoundException;
//import java.util.Collections;
//import java.util.List;
//import java.util.stream.Collectors;
//
//import javax.servlet.FilterChain;
//import javax.servlet.ServletException;
//import javax.servlet.ServletRequest;
//import javax.servlet.ServletResponse;
//import javax.servlet.http.HttpServletRequest;
//
//import org.springframework.beans.factory.annotation.Autowired;
//import org.springframework.security.authentication.BadCredentialsException;
//import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
//import org.springframework.security.core.AuthenticationException;
//import org.springframework.security.core.context.SecurityContextHolder;
//import org.springframework.security.web.util.matcher.AntPathRequestMatcher;
//import org.springframework.security.web.util.matcher.RequestMatcher;
//import org.springframework.web.filter.GenericFilterBean;
//
//import com.google.common.base.Strings;
//
//import xdean.jex.log.Logable;
//import xdean.mini.boardgame.server.security.OpenIdAuthProvider;
//
//public class OpenIdAuthFilter extends GenericFilterBean implements Logable {
//
//  private RequestMatcher requiresAuthenticationRequestMatcher;
//
//  @Autowired(required = false)
//  List<OpenIdAuthProvider> providers = Collections.emptyList();
//
//  public OpenIdAuthFilter() {
//    this(new AntPathRequestMatcher("/login"));
//  }
//
//  public OpenIdAuthFilter(RequestMatcher requiresAuthenticationRequestMatcher) {
//    this.requiresAuthenticationRequestMatcher = requiresAuthenticationRequestMatcher;
//  }
//
//  @Override
//  public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain) throws IOException, ServletException {
//    if (!requiresAuthenticationRequestMatcher.matches((HttpServletRequest) request)) {
//      chain.doFilter(request, response);
//      return;
//    }
//    String type = request.getParameter("type");
//    if (!"openid".equals(type)) {
//      chain.doFilter(request, response);
//      return;
//    }
//    String provider = Strings.nullToEmpty(request.getParameter("provider"));
//    String token = Strings.nullToEmpty(request.getParameter("token"));
//
//    try {
//      List<OpenIdAuthProvider> findProviders = providers.stream().filter(p -> p.name().equals(provider))
//          .collect(Collectors.toList());
//      if (findProviders.isEmpty()) {
//        throw new ProviderNotFoundException("There is no provider support: " + provider);
//      }
//      for (int i = 0; i < findProviders.size(); i++) {
//        OpenIdAuthProvider p = findProviders.get(i);
//        try {
//          String result = p.attemptAuthentication(token);
//          if (result != null) {
//            SecurityContextHolder.getContext().setAuthentication(new UsernamePasswordAuthenticationToken(result, null));
//          }
//        } catch (AuthenticationException e) {
//          trace().log("Fail to authenticate: " + token, e);
//        }
//      }
//      throw new BadCredentialsException("Bad Credentials");
//    } catch (AuthenticationException e) {
//      unsuccessfulAuthentication(e);
//    }
//  }
//
//  protected void unsuccessfulAuthentication(AuthenticationException failed) throws IOException, ServletException {
//    SecurityContextHolder.clearContext();
//    if (logger.isDebugEnabled()) {
//      logger.debug("Authentication request failed: " + failed.toString(), failed);
//      logger.debug("Updated SecurityContextHolder to contain null Authentication");
//    }
//  }
//}
