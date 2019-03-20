package xdean.mini.boardgame.server.annotation.processor;

import java.io.IOException;
import java.io.PrintStream;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

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
import xdean.annotation.processor.toolkit.AssertException;
import xdean.annotation.processor.toolkit.ElementUtil;
import xdean.annotation.processor.toolkit.XAbstractProcessor;
import xdean.annotation.processor.toolkit.annotation.SupportedAnnotation;
import xdean.jex.util.string.StringUtil;
import xdean.mini.boardgame.server.annotation.Attr;
import xdean.mini.boardgame.server.annotation.Payload;
import xdean.mini.boardgame.server.annotation.Side;
import xdean.mini.boardgame.server.annotation.Topic;
import xdean.mini.boardgame.server.annotation.TopicDoc;
import xdean.mini.boardgame.server.annotation.processor.model.SocketAttr;
import xdean.mini.boardgame.server.annotation.processor.model.SocketPayload;
import xdean.mini.boardgame.server.annotation.processor.model.SocketSide;
import xdean.mini.boardgame.server.annotation.processor.model.SocketSide.SocketSideBuilder;
import xdean.mini.boardgame.server.annotation.processor.model.SocketTopic;
import xdean.mini.boardgame.server.annotation.processor.model.SocketTopicGroup;

@AutoService(Processor.class)
@SupportedAnnotation({ TopicDoc.class, Attr.class })
@SupportedSourceVersion(SourceVersion.RELEASE_8)
public class SocketDocumentProcessor extends XAbstractProcessor {

  private final Map<String, SocketAttr> attrs = new HashMap<>();
  private final Map<TopicDoc, SocketTopicGroup> groups = new HashMap<>();

  @Override
  public boolean processActual(Set<? extends TypeElement> annotations, RoundEnvironment roundEnv) throws AssertException {
    if (roundEnv.processingOver()) {
      groups.forEach((anno, g) -> {
        String document = generateDocument(g);
        try {
          String path = anno.path();
          if (path.startsWith("/")) {
            path = path.substring(1);
          }
          FileObject file = filer.createResource(StandardLocation.CLASS_OUTPUT, "", path);
          PrintStream ps = new PrintStream(file.openOutputStream());
          ps.print(document);
          ps.close();
        } catch (IOException e) {
          e.printStackTrace();
          throw new Error(e);
        }
      });
      return false;
    }
    Set<VariableElement> attrs = ElementFilter.fieldsIn(roundEnv.getElementsAnnotatedWith(Attr.class));
    attrs.forEach(e -> processAttr(e, e.getAnnotation(Attr.class)));

    Set<TypeElement> topicDocs = ElementFilter.typesIn(roundEnv.getElementsAnnotatedWith(TopicDoc.class));
    topicDocs.forEach(d -> {
      List<VariableElement> topics = ElementFilter.fieldsIn(
          d.getEnclosedElements().stream().filter(e -> e.getAnnotation(Topic.class) != null).collect(Collectors.toList()));
      SocketTopicGroup group = new SocketTopicGroup("Topics", 0);
      topics.forEach(e -> group.add(processTopic(e)));
      groups.put(d.getAnnotation(TopicDoc.class), group);
    });
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

  private SocketTopic processTopic(VariableElement e) {
    assertNonNull(e.getConstantValue());
    Topic anno = e.getAnnotation(Topic.class);
    String topic = e.getConstantValue().toString();
    return SocketTopic.builder()
        .topic(topic)
        .category(anno.category())
        .fromClient(processSide(anno.fromClient()))
        .fromServer(processSide(anno.fromServer()))
        .build();
  }

  private SocketSide processSide(Side anno) {
    if (anno.disable()) {
      return null;
    }
    SocketSideBuilder sideBuilder = SocketSide.builder()
        .desc(description(anno.desc()));
    Arrays.stream(anno.attr()).forEach(attr -> sideBuilder.attr(processAttr(null, attr)));
    sideBuilder.fromServer(false);
    sideBuilder.payload(processPayload(anno.payload()));
    SocketSide side = sideBuilder.build();
    return side;
  }

  private SocketPayload processPayload(Payload payload) {
    TypeMirror type = ElementUtil.getAnnotationClassValue(elements, payload, p -> p.type());
    boolean noType = type.toString().equals(void.class.getName());
    if (noType && payload.desc().isEmpty()) {
      return null;
    } else {
      return SocketPayload.builder()
          .type(noType ? null : type)
          .desc(description(payload.desc()))
          .build();
    }
  }

  private String generateDocument(SocketTopicGroup group) {
    StringBuilder sb = new StringBuilder();
    sb.append("# Socket Topics").append("\n");
    sb.append("---\n\n");
    sb.append("Table of Contents\n---\n\n");
    group.forEach((e, level) -> {
      sb.append(StringUtil.repeat("  ", level - 1)).append("- ");
      String name = e.unify(g -> g.getName(), t -> t.getTopic());
      sb.append(String.format("[%s](%s)", name, titleToLink(name))).append("\n");
    });

    group.forEach((e, level) -> {
      e.exec(g -> {
        sb.append(String.format("<a name=\"%s\"></a>\n\n", g.getName()));
        sb.append(StringUtil.repeat("#", level)).append(" ").append(g.getName()).append("\n---\n\n");
      }, t -> {
        sb.append(String.format("<a name=\"%s\"></a>\n\n", t.getTopic()));
        sb.append(StringUtil.repeat("#", level)).append(" ").append(t.getTopic()).append("\n---\n\n");
        sb.append("Topic: `").append(t.getTopic()).append("`\n\n");
        sb.append(formatSide(t.getFromServer()));
        sb.append(formatSide(t.getFromClient()));
      });
    });
    return sb.toString();
  }

  private String titleToLink(String title) {
    return "#" + title.toLowerCase().replace(" ", "-").replace("`", "").replace("*", "");
  }

  private String formatSide(SocketSide side) {
    if (side == null) {
      return "";
    }
    StringBuilder sb = new StringBuilder();
    sb.append(new BoldText(side.isFromServer() ? "From Server" : "From Client")).append("\n\n");
    sb.append(side.getDesc()).append("\n\n");
    Table.Builder table = new Builder()
        .withAlignments(Table.ALIGN_CENTER)
        .addRow("Name", "Type", "Description");
    side.getAttrs().forEach(attr -> table.addRow(attr.getKey(), code(attr.getType().toString()), attr.getDesc()));
    if (side.getPayload() != null) {
      TypeMirror type = side.getPayload().getType();
      table.addRow(new BoldText("Payload"), code(type == null ? "Undefined" : type.toString()), side.getPayload().getDesc());
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
