package xdean.mini.boardgame.server.service;

import java.io.IOException;
import java.io.UnsupportedEncodingException;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.web.servlet.DispatcherServlet;
import org.springframework.web.servlet.HandlerExecutionChain;
import org.springframework.web.util.ContentCachingRequestWrapper;
import org.springframework.web.util.ContentCachingResponseWrapper;
import org.springframework.web.util.WebUtils;

import com.google.common.base.MoreObjects;
import com.google.common.base.MoreObjects.ToStringHelper;

import xdean.jex.log.Log;
import xdean.jex.log.LogFactory;
import xdean.jex.log.Logable;

public class LoggableDispatcherServlet extends DispatcherServlet implements Logable {

  private static final Log LOG = LogFactory.from("xdean.printResponse");

  @Override
  public Log log() {
    return LOG;
  }

  @Override
  protected void doDispatch(HttpServletRequest request, HttpServletResponse response) throws Exception {
    if (!(request instanceof ContentCachingRequestWrapper)) {
      request = new ContentCachingRequestWrapper(request);
    }
    if (!(response instanceof ContentCachingResponseWrapper)) {
      response = new ContentCachingResponseWrapper(response);
    }
    HandlerExecutionChain handler = getHandler(request);

    try {
      super.doDispatch(request, response);
    } finally {
      log(request, response, handler);
      updateResponse(response);
    }
  }

  private void log(HttpServletRequest request, HttpServletResponse response, HandlerExecutionChain handler) {
    if (isTraceEnable()) {
      String contentType = response.getContentType();
      ToStringHelper tsh = MoreObjects.toStringHelper("Http request and response")
          .add("status", response.getStatus())
          .add("method", request.getMethod())
          .add("path", request.getRequestURI())
          .add("client-ip", request.getRemoteAddr())
          .add("java-method", handler.toString())
          .add("content-type", contentType);
      if (contentType != null && (contentType.contains("text") || contentType.contains("json") || contentType.contains("xml"))) {
        tsh.add("response", getResponsePayload(response));
      }
      trace(tsh.toString());
    }
  }

  private String getResponsePayload(HttpServletResponse response) {
    ContentCachingResponseWrapper wrapper = WebUtils.getNativeResponse(response, ContentCachingResponseWrapper.class);
    if (wrapper != null) {
      byte[] buf = wrapper.getContentAsByteArray();
      if (buf.length > 0) {
        int length = Math.min(buf.length, 5120);
        try {
          return new String(buf, 0, length, wrapper.getCharacterEncoding());
        } catch (UnsupportedEncodingException ex) {
          // NOOP
        }
      }
    }
    return "[unknown]";
  }

  private void updateResponse(HttpServletResponse response) throws IOException {
    ContentCachingResponseWrapper responseWrapper = WebUtils.getNativeResponse(response, ContentCachingResponseWrapper.class);
    responseWrapper.copyBodyToResponse();
  }
}