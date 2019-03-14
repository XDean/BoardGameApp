package xdean.mini.boardgame.server.mybatis.mapper;

import xdean.mybatis.extension.model.Column;

public interface BaseMapper {
  int countDistinct(Column column);

  boolean exist(Column column, String value);
}
