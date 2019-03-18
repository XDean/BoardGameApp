package xdean.mini.boardgame.server.annotation.processor;

import java.util.Set;

import javax.annotation.processing.Processor;
import javax.annotation.processing.RoundEnvironment;
import javax.annotation.processing.SupportedSourceVersion;
import javax.lang.model.SourceVersion;
import javax.lang.model.element.Element;
import javax.lang.model.element.TypeElement;

import com.google.auto.service.AutoService;

import xdean.annotation.processor.toolkit.AssertException;
import xdean.annotation.processor.toolkit.XAbstractProcessor;
import xdean.annotation.processor.toolkit.annotation.SupportedAnnotation;
import xdean.mini.boardgame.server.annotation.FromClient;
import xdean.mini.boardgame.server.annotation.FromServer;

@AutoService(Processor.class)
@SupportedAnnotation({ FromServer.class, FromClient.class })
@SupportedSourceVersion(SourceVersion.RELEASE_8)
public class SocketDocumentProcessor extends XAbstractProcessor {

  @Override
  public boolean processActual(Set<? extends TypeElement> annotations, RoundEnvironment roundEnv) throws AssertException {
    if (roundEnv.processingOver()) {
      return false;
    }
    Set<? extends Element> fromServers = roundEnv.getElementsAnnotatedWith(FromServer.class);
    Set<? extends Element> fromClients = roundEnv.getElementsAnnotatedWith(FromClient.class);

    return true;
  }
}
