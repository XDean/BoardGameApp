package xdean.mini.boardgame.server.mybatis.mapper;

import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mybatis.extension.MyBatisSQL;

public class UserMapperBuilder {
  String save(UserProfileEntity build) {
    return MyBatisSQL.create()
        .toString();
  }

  String findByUsername(String username) {
    return null;
  }
}
