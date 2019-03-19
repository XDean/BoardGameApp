package xdean.mini.boardgame.server.annotation.processor.model;

import java.util.ArrayList;
import java.util.List;
import java.util.function.BiConsumer;

import lombok.Data;
import xdean.jex.extra.collection.Either;

@Data
public class SocketTopicGroup {
  final String name;

  final int level;

  List<SocketTopicGroup> groups = new ArrayList<>();

  List<SocketTopic> topics = new ArrayList<>();

  public void forEach(BiConsumer<Either<SocketTopicGroup, SocketTopic>, Integer> run) {
    groups.forEach(g -> {
      run.accept(Either.left(g), level + 1);
      g.forEach(run);
    });
    topics.forEach(t -> run.accept(Either.right(t), level + 1));
  }

  public SocketTopicGroup add(SocketTopic topic) {
    if (topic.getCategory().length > level) {
      String category = topic.getCategory()[level];
      groups.stream()
          .filter(g -> g.name.equals(category))
          .findAny()
          .orElseGet(() -> {
            SocketTopicGroup group = new SocketTopicGroup(category, level + 1);
            groups.add(group);
            return group;
          })
          .add(topic);
    } else {
      topics.add(topic);
    }
    return this;
  }
}
