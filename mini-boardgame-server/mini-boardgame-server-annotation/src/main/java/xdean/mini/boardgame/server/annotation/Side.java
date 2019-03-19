package xdean.mini.boardgame.server.annotation;

import static java.lang.annotation.ElementType.ANNOTATION_TYPE;
import static java.lang.annotation.RetentionPolicy.RUNTIME;

import java.lang.annotation.Retention;
import java.lang.annotation.Target;

import javax.lang.model.element.Modifier;

import xdean.deannotation.checker.CheckField;
import xdean.deannotation.checker.CheckModifier;
import xdean.deannotation.checker.CheckType;

@Retention(RUNTIME)
@Target(ANNOTATION_TYPE)
@CheckField(
    type = @CheckType(String.class),
    modifier = @CheckModifier(require = { Modifier.STATIC, Modifier.FINAL }))
public @interface Side {
  Attr[] attr() default {};

  Payload payload() default @Payload;

  String desc() default "";

  boolean disable() default false;
}
