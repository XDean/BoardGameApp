package xdean.mini.boardgame.server.mybatis.mapper;

import static xdean.mybatis.extension.SqlUtil.equal;

import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mini.boardgame.server.mybatis.Tables;
import xdean.mybatis.extension.MyBatisSQL;
import xdean.mybatis.extension.SqlUtil;

public class UserMapperBuilder implements Tables {
  String save(UserProfileEntity e) {
    return MyBatisSQL.create()
        .INSERT_INTO(ProfileTable.table)
        .VALUES(ProfileTable.id.fullName, Integer.toString(e.getId()))
        .VALUES(ProfileTable.avatarUrl.fullName, e.getAvatarUrl())
        .VALUES(ProfileTable.nickname.fullName, e.getNickname())
        .VALUES(ProfileTable.male.fullName, Boolean.toString(e.isMale()))
        .ON_DUPLICATE_KEY_UPDATE(ProfileTable.avatarUrl, ProfileTable.nickname, ProfileTable.male)
        .toString();
  }

  String findByUsername(String username) {
    return MyBatisSQL.create()
        .SELECT_ALL(UserTable.table)
        .SELECT_ALL(ProfileTable.table)
        .FROM(UserTable.table)
        .INNER_JOIN(ProfileTable.table, equal(UserTable.id, ProfileTable.id))
        .WHERE(SqlUtil.equal(UserTable.username.fullName, username))
        .toString();
  }

  String findAuthorities(int id) {
    return MyBatisSQL.create()
        .SELECT(AuthorityTable.authority)
        .FROM(AuthorityTable.table)
        .WHERE(equal(AuthorityTable.id.fullName, id))
        .toString();
  }
}
