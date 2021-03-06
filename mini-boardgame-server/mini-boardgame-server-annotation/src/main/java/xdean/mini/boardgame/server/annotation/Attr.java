package xdean.mini.boardgame.server.annotation;

import static java.lang.annotation.ElementType.FIELD;
import static java.lang.annotation.RetentionPolicy.RUNTIME;

import java.lang.annotation.Retention;
import java.lang.annotation.Target;

import javax.lang.model.element.Modifier;

import xdean.deannotation.checker.CheckField;
import xdean.deannotation.checker.CheckModifier;
import xdean.deannotation.checker.CheckType;

@Retention(RUNTIME)
@Target(FIELD)
@CheckField(
    type = @CheckType(String.class),
    modifier = @CheckModifier(require = { Modifier.STATIC, Modifier.FINAL }))
public @interface Attr {
  String value/* id */() default "";

  Class<?> type() default void.class;

  String desc() default "";
}
