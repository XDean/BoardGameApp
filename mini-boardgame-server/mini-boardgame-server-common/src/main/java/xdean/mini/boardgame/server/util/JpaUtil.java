package xdean.mini.boardgame.server.util;

import java.util.Optional;
import java.util.function.Function;

import org.springframework.data.jpa.repository.JpaRepository;

public enum JpaUtil {
  ;
  public static <T, ID> T findOrCreate(JpaRepository<T, ID> repo, ID id, Function<ID, T> creator) {
    Optional<T> find = repo.findById(id);
    if (find.isPresent()) {
      return find.get();
    } else {
      T create = creator.apply(id);
      return repo.save(create);
    }
  }
}
