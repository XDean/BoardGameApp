package xdean.mini.boardgame.server.annotation;

import static java.lang.annotation.ElementType.FIELD;
import static java.lang.annotation.RetentionPolicy.RUNTIME;

import java.lang.annotation.Retention;
import java.lang.annotation.Target;

@Retention(RUNTIME)
@Target(FIELD)
public @interface FromServer {
  Attr[] attr() default {};

  Payload payload() default @Payload;

  String desc() default "";
}