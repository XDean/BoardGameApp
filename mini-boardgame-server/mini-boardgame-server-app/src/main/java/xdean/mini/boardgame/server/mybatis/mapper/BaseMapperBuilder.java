package xdean.mini.boardgame.server.mybatis.mapper;

import static xdean.mybatis.extension.SqlUtil.count;
import static xdean.mybatis.extension.SqlUtil.distinct;
import static xdean.mybatis.extension.SqlUtil.equal;

import xdean.mybatis.extension.MyBatisSQL;
import xdean.mybatis.extension.model.Column;

public class BaseMapperBuilder {

  public String countDistinct(Column column) {
    return MyBatisSQL.create()
        .SELECT(count(distinct(column.fullName)))
        .FROM(column.table)
        .toString();
  }

  public String exist(Column column, String value) {
    return MyBatisSQL.create()
        .SELECT(count())
        .FROM(MyBatisSQL.create()
            .SELECT(column)
            .FROM(column.table)
            .WHERE(equal(column.fullName, value))
            .LIMIT(1)
            .toString())
        .toString();
  }
}
