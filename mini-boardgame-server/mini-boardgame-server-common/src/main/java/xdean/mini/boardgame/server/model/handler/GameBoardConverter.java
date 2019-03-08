package xdean.mini.boardgame.server.model.handler;

import java.io.IOException;

import javax.persistence.AttributeConverter;
import javax.persistence.Converter;

import com.fasterxml.jackson.annotation.JsonAutoDetect.Visibility;
import com.fasterxml.jackson.annotation.PropertyAccessor;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.common.base.Strings;
import com.google.common.collect.ImmutableMap;

import xdean.jex.log.Logable;
import xdean.mini.boardgame.server.model.GameBoard;

@Converter(autoApply = true)
public class GameBoardConverter implements AttributeConverter<GameBoard, String>, Logable {
  ObjectMapper objectMapper = new ObjectMapper();

  public GameBoardConverter() {
    objectMapper.setVisibility(PropertyAccessor.ALL, Visibility.NONE);
    objectMapper.setVisibility(PropertyAccessor.FIELD, Visibility.ANY);
  }

  @Override
  public String convertToDatabaseColumn(GameBoard attribute) {
    if (attribute == null) {
      return "";
    }
    try {
      String clz = attribute.getClass().getName();
      String value = objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(attribute);
      return objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(ImmutableMap.of(
          "class", clz,
          "value", value));
    } catch (JsonProcessingException e) {
      debug("Fail to serialize game board: " + attribute, e);
      return null;
    }
  }

  @Override
  @SuppressWarnings("unchecked")
  public GameBoard convertToEntityAttribute(String dbData) {
    if (Strings.isNullOrEmpty(dbData)) {
      return null;
    }
    try {
      JsonNode map = objectMapper.readValue(dbData, JsonNode.class);
      Class<? extends GameBoard> clz = (Class<? extends GameBoard>) Class.forName(map.get("class").textValue());
      String valueStr = map.get("value").textValue();
      return objectMapper.readValue(valueStr, clz);
    } catch (IOException | ClassNotFoundException e) {
      debug("Fail to deserialize game board: " + dbData, e);
      return null;
    }
  }
}
