package xdean.mini.boardgame.server.mybatis.mapper;

import xdean.mini.boardgame.server.model.Tables;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mybatis.extension.MyBatisSQL;
import xdean.mybatis.extension.SqlUtil;

public class UserMapperBuilder implements Tables {
  String save(UserProfileEntity e) {
    return MyBatisSQL.create()
        .INSERT_INTO(ProfileTable.table)
        .VALUES(ProfileTable.id.fullName, Integer.toString(e.getUserId()))
        .VALUES(ProfileTable.avatarUrl.fullName, e.getAvatarUrl())
        .VALUES(ProfileTable.nickname.fullName, e.getNickname())
        .VALUES(ProfileTable.male.fullName, Boolean.toString(e.isMale()))
        .ON_DUPLICATE_KEY_UPDATE(ProfileTable.avatarUrl, ProfileTable.nickname, ProfileTable.male)
        .toString();
  }

  String findByUsername(String username) {
    return MyBatisSQL.create()
        .SELECT_FROM(UserTable.table)
        .WHERE(SqlUtil.equal(UserTable.username.fullName, username))
        .toString();
  }
}
