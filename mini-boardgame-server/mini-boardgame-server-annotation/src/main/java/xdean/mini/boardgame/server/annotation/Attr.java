package xdean.mini.boardgame.server.annotation;

import static java.lang.annotation.ElementType.*;
import static java.lang.annotation.RetentionPolicy.RUNTIME;

import java.lang.annotation.Retention;
import java.lang.annotation.Target;

@Retention(RUNTIME)
@Target(FIELD)
public @interface Attr {
  String value/* id */() default "";

  Class<?> type() default void.class;

  String desc() default "";
}
