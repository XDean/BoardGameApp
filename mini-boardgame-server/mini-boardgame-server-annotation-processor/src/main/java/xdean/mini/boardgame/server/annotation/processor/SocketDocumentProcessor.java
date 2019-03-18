package xdean.mini.boardgame.server.annotation.processor;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;

import javax.annotation.CheckForNull;
import javax.annotation.processing.Processor;
import javax.annotation.processing.RoundEnvironment;
import javax.annotation.processing.SupportedSourceVersion;
import javax.lang.model.SourceVersion;
import javax.lang.model.element.TypeElement;
import javax.lang.model.element.VariableElement;
import javax.lang.model.util.ElementFilter;

import com.google.auto.service.AutoService;

import xdean.annotation.processor.toolkit.AssertException;
import xdean.annotation.processor.toolkit.ElementUtil;
import xdean.annotation.processor.toolkit.XAbstractProcessor;
import xdean.annotation.processor.toolkit.annotation.SupportedAnnotation;
import xdean.mini.boardgame.server.annotation.Attr;
import xdean.mini.boardgame.server.annotation.FromClient;
import xdean.mini.boardgame.server.annotation.FromServer;
import xdean.mini.boardgame.server.annotation.Payload;
import xdean.mini.boardgame.server.annotation.processor.model.SocketAttr;
import xdean.mini.boardgame.server.annotation.processor.model.SocketDescription.SocketDescriptionBuilder;
import xdean.mini.boardgame.server.annotation.processor.model.SocketPayload;
import xdean.mini.boardgame.server.annotation.processor.model.SocketSide;
import xdean.mini.boardgame.server.annotation.processor.model.SocketSide.SocketSideBuilder;

@AutoService(Processor.class)
@SupportedAnnotation({ FromServer.class, FromClient.class, Attr.class })
@SupportedSourceVersion(SourceVersion.RELEASE_8)
public class SocketDocumentProcessor extends XAbstractProcessor {

  private final Map<String, SocketAttr> attrs = new HashMap<>();
  private final Map<String, SocketDescriptionBuilder> descBuilders = new HashMap<>();

  @Override
  public boolean processActual(Set<? extends TypeElement> annotations, RoundEnvironment roundEnv) throws AssertException {
    if (roundEnv.processingOver()) {
      return false;
    }
    Set<VariableElement> attrs = ElementFilter.fieldsIn(roundEnv.getElementsAnnotatedWith(Attr.class));
    attrs.forEach(e -> processAttr(e, e.getAnnotation(Attr.class)));
    Set<VariableElement> fromServers = ElementFilter.fieldsIn(roundEnv.getElementsAnnotatedWith(FromServer.class));
    Set<VariableElement> fromClients = ElementFilter.fieldsIn(roundEnv.getElementsAnnotatedWith(FromClient.class));
    fromServers.forEach(e -> processFromServerEvent(e));
    return true;
  }

  private SocketAttr processAttr(@CheckForNull VariableElement e, Attr anno) {
    String id;
    if (!anno.value().isEmpty()) {
      id = anno.value();
    } else {
      assertNonNull(e).log("@Attr ref must have `value`:" + anno);
      if (e.getConstantValue() != null) {
        id = e.getConstantValue().toString();
      } else {
        id = e.getSimpleName().toString();
      }
    }
    String desc = description(anno.desc());
    // it's a reference
    if (ElementUtil.getAnnotationClassValue(elements, anno, a -> a.type()).toString().equals(void.class.toString())) {
      SocketAttr cache = attrs.get(id);
      assertNonNull(cache).log("No such @Attr to reference", e);
      if (anno.desc().isEmpty()) {
        return cache;
      } else {
        return cache.toBuilder().desc(desc).build();
      }
    }
    SocketAttr attr = SocketAttr.builder()
        .id(id)
        .desc(desc)
        .type(ElementUtil.getAnnotationClassValue(elements, anno, a -> a.type()))
        .build();
    attrs.put(id, attr);
    return attr;
  }

  private void processFromServerEvent(VariableElement e) {
    assertNonNull(e.getConstantValue());

    FromServer anno = e.getAnnotation(FromServer.class);
    SocketSideBuilder sideBuilder = SocketSide.builder()
        .desc(description(anno.desc()));
    Arrays.stream(anno.attr()).forEach(attr -> sideBuilder.attr(processAttr(null, attr)));
    sideBuilder.fromServer(true);
    sideBuilder.payload(processPayload(anno.payload()));
    SocketSide side = sideBuilder.build();
  }

  private SocketPayload processPayload(Payload payload) {
    return SocketPayload.builder()
        .desc(description(payload.desc()))
        .type(ElementUtil.getAnnotationClassValue(elements, payload, p -> p.value()))
        .build();
  }

  private String description(String desc) {
    return desc.isEmpty() ? "No Description" : desc;
  }
}
