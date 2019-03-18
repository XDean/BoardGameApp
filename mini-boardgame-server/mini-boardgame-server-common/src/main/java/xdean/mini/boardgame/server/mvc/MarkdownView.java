package xdean.mini.boardgame.server.mvc;

import java.io.IOException;
import java.io.PrintWriter;
import java.net.URISyntaxException;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.List;
import java.util.Locale;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.commonmark.Extension;
import org.commonmark.ext.gfm.tables.TablesExtension;
import org.commonmark.node.Node;
import org.commonmark.parser.Parser;
import org.commonmark.renderer.html.HtmlRenderer;
import org.springframework.web.servlet.view.AbstractTemplateView;

public class MarkdownView extends AbstractTemplateView {
  @Override
  public boolean checkResource(Locale locale) throws Exception {
    String templatePath = "static/" + getUrl();
    URL templateUrl = MarkdownView.class.getClassLoader().getResource(templatePath);
    return templateUrl != null;
  }

  @Override
  protected void renderMergedTemplateModel(
      Map<String, Object> model,
      HttpServletRequest request,
      HttpServletResponse response) throws Exception {

    PrintWriter writer = response.getWriter();
    writer.append("<html>\n" +
        "<head>\n" +
        "  <meta charset=\"utf-8\"/>\n" +
        "  <link rel=\"stylesheet\" type=\"text/css\" href=\"/css/markdown.css\"/>\n" +
        "</head><body>");
    writer.append(getHtmlFromMarkdown());
    writer.append("</body></html>");
  }

  private String getHtmlFromMarkdown() throws URISyntaxException, IOException {
    String templatePath = "static/" + getUrl();
    URL templateUrl = MarkdownView.class.getClassLoader().getResource(templatePath);
    Path path = Paths.get(templateUrl.toURI());

    String markdown = new String(Files.readAllBytes(path));

    List<Extension> extensions = Arrays.asList(TablesExtension.create());
    Parser parser = Parser.builder().extensions(extensions).build();
    Node document = parser.parse(markdown);
    HtmlRenderer renderer = HtmlRenderer.builder().extensions(extensions).build();
    return renderer.render(document);
  }
}