package xdean.mini.boardgame.server.model.handler;

import java.io.IOException;

import com.fasterxml.jackson.annotation.JsonAutoDetect.Visibility;
import com.fasterxml.jackson.annotation.PropertyAccessor;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.common.base.Strings;
import com.google.common.collect.ImmutableMap;

import xdean.jex.log.Logable;
import xdean.mybatis.extension.resultmap.StringParseHandler;

public class ObjectJsonConverter implements StringParseHandler<Object>, Logable {
  ObjectMapper objectMapper = new ObjectMapper();

  public ObjectJsonConverter() {
    objectMapper.setVisibility(PropertyAccessor.ALL, Visibility.NONE);
    objectMapper.setVisibility(PropertyAccessor.FIELD, Visibility.ANY);
  }

  @Override
  public String toString(Object board) {
    if (board == null) {
      return "";
    }
    try {
      String clz = board.getClass().getName();
      return objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(ImmutableMap.of(
          "class", clz,
          "value", board));
    } catch (JsonProcessingException e) {
      debug("Fail to serialize game board: " + board, e);
      return null;
    }
  }

  @Override
  public Object parse(String dbData) {
    if (Strings.isNullOrEmpty(dbData)) {
      return null;
    }
    try {
      JsonNode map = objectMapper.readValue(dbData, JsonNode.class);
      Class<?> clz = Class.forName(map.get("class").textValue());
      return objectMapper.treeToValue(map.get("value"), clz);
    } catch (IOException | ClassNotFoundException e) {
      debug("Fail to deserialize game board: " + dbData, e);
      return null;
    }
  }
}
