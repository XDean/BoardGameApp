package xdean.mini.boardgame.server.annotation.processor;

import java.io.IOException;
import java.io.PrintStream;
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
import javax.lang.model.type.TypeMirror;
import javax.lang.model.util.ElementFilter;
import javax.tools.FileObject;
import javax.tools.StandardLocation;

import com.google.auto.service.AutoService;

import net.steppschuh.markdowngenerator.table.Table;
import net.steppschuh.markdowngenerator.table.Table.Builder;
import net.steppschuh.markdowngenerator.text.emphasis.BoldText;
import net.steppschuh.markdowngenerator.text.heading.Heading;
import xdean.annotation.processor.toolkit.AssertException;
import xdean.annotation.processor.toolkit.ElementUtil;
import xdean.annotation.processor.toolkit.XAbstractProcessor;
import xdean.annotation.processor.toolkit.annotation.SupportedAnnotation;
import xdean.mini.boardgame.server.annotation.Attr;
import xdean.mini.boardgame.server.annotation.FromClient;
import xdean.mini.boardgame.server.annotation.FromServer;
import xdean.mini.boardgame.server.annotation.Payload;
import xdean.mini.boardgame.server.annotation.processor.model.SocketAttr;
import xdean.mini.boardgame.server.annotation.processor.model.SocketDescription;
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
      String document = generateDocument();
      try {
        FileObject file = filer.createResource(StandardLocation.CLASS_OUTPUT, "", "static/doc/socket.md");
        PrintStream ps = new PrintStream(file.openOutputStream());
        ps.print(document);
        ps.close();
      } catch (IOException e) {
        e.printStackTrace();
        throw new Error(e);
      }
      return false;
    }
    Set<VariableElement> attrs = ElementFilter.fieldsIn(roundEnv.getElementsAnnotatedWith(Attr.class));
    attrs.forEach(e -> processAttr(e, e.getAnnotation(Attr.class)));
    Set<VariableElement> fromServers = ElementFilter.fieldsIn(roundEnv.getElementsAnnotatedWith(FromServer.class));
    Set<VariableElement> fromClients = ElementFilter.fieldsIn(roundEnv.getElementsAnnotatedWith(FromClient.class));
    fromServers.forEach(e -> processFromServerEvent(e));
    fromClients.forEach(e -> processFromClientEvent(e));
    return true;
  }

  private SocketAttr processAttr(@CheckForNull VariableElement e, Attr anno) {
    String id;
    if (!anno.value().isEmpty()) {
      id = anno.value();
    } else {
      assertNonNull(e).log("@Attr ref must have `value`:" + anno);
      id = e.getConstantValue().toString();
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
        .key(id)
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

    String topic = e.getConstantValue().toString();
    descBuilders.computeIfAbsent(topic, t -> SocketDescription.builder().topic(t))
        .fromServer(side);
  }

  private void processFromClientEvent(VariableElement e) {
    assertNonNull(e.getConstantValue());

    FromClient anno = e.getAnnotation(FromClient.class);
    SocketSideBuilder sideBuilder = SocketSide.builder()
        .desc(description(anno.desc()));
    Arrays.stream(anno.attr()).forEach(attr -> sideBuilder.attr(processAttr(null, attr)));
    sideBuilder.fromServer(false);
    sideBuilder.payload(processPayload(anno.payload()));
    SocketSide side = sideBuilder.build();

    String topic = e.getConstantValue().toString();
    descBuilders.computeIfAbsent(topic, t -> SocketDescription.builder().topic(t))
        .fromClient(side);
  }

  private SocketPayload processPayload(Payload payload) {
    TypeMirror type = ElementUtil.getAnnotationClassValue(elements, payload, p -> p.value());
    if (type.toString().equals(void.class.getName()) && payload.desc().isEmpty()) {
      return null;
    } else {
      return SocketPayload.builder()
          .type(type)
          .desc(description(payload.desc()))
          .build();
    }
  }

  private String generateDocument() {
    StringBuilder sb = new StringBuilder();
    sb.append(new Heading("Socket Topics", 1)).append("\n");
    sb.append("---\n\n");
    descBuilders.values().forEach(builder -> {
      SocketDescription desc = builder.build();
      sb.append(new Heading(desc.getTopic(), 2)).append("\n");
      sb.append("---\n\n");
      sb.append(formatSide(desc.getFromServer()));
      sb.append(formatSide(desc.getFromClient()));
    });
    return sb.toString();
  }

  private String formatSide(SocketSide side) {
    if (side == null) {
      return "";
    }
    StringBuilder sb = new StringBuilder();
    sb.append(new Heading(side.isFromServer() ? "From Server" : "From Client", 3)).append("\n\n");
    sb.append(side.getDesc()).append("\n\n");
    Table.Builder table = new Builder()
        .withAlignments(Table.ALIGN_CENTER)
        .addRow("Name", "Type", "Description");
    side.getAttrs().forEach(attr -> table.addRow(attr.getKey(), code(attr.getType().toString()), attr.getDesc()));
    if (side.getPayload() != null) {
      table.addRow(new BoldText("Payload"), code(side.getPayload().getType().toString()), side.getPayload().getDesc());
    }
    if (!side.getAttrs().isEmpty() || side.getPayload() != null) {
      sb.append(table.build()).append("\n\n");
    }
    return sb.toString();
  }

  private String description(String desc) {
    return desc.isEmpty() ? "No Description" : desc;
  }

  private String code(String code) {
    return "`" + code + "`";
  }
}
