package xdean.mini.boardgame.server.mybatis.mapper;

import static xdean.mybatis.extension.SqlUtil.equal;
import static xdean.mybatis.extension.SqlUtil.together;
import static xdean.mybatis.extension.SqlUtil.wrapString;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mini.boardgame.server.mybatis.Tables;
import xdean.mybatis.extension.MyBatisSQL;
import xdean.mybatis.extension.MyBatisSQL.MSB;
import xdean.mybatis.extension.SqlUtil;

public class UserMapperBuilder implements Tables {

  public String findByUsername(String username) {
    return MyBatisSQL.create()
        .SELECT_ALL(UserTable.table)
        .SELECT_ALL(ProfileTable.table)
        .FROM(UserTable.table)
        .INNER_JOIN(ProfileTable.table, equal(UserTable.id, ProfileTable.id))
        .WHERE(SqlUtil.equal(UserTable.username.fullName, username))
        .toString();
  }

  public String findAuthorities(int id) {
    return MyBatisSQL.create()
        .SELECT(AuthorityTable.authority)
        .FROM(AuthorityTable.table)
        .WHERE(equal(AuthorityTable.id.fullName, id))
        .toString();
  }

  public String save(UserProfileEntity e) {
    return MyBatisSQL.create()
        .INSERT_INTO(ProfileTable.table)
        .VALUES(ProfileTable.id.fullName, Integer.toString(e.getId()))
        .VALUES(ProfileTable.avatarUrl.fullName, e.getAvatarUrl())
        .VALUES(ProfileTable.nickname.fullName, e.getNickname())
        .VALUES(ProfileTable.male.fullName, Boolean.toString(e.isMale()))
        .ON_DUPLICATE_KEY_UPDATE(ProfileTable.avatarUrl, ProfileTable.nickname, ProfileTable.male)
        .toString();
  }

  public String save(UserEntity user) {
    return together(
        MyBatisSQL.create()
            .INSERT_INTO(UserTable.table)
            .VALUES(UserTable.username.fullName, user.getUsername())
            .VALUES(UserTable.password.fullName, user.getPassword())
            .VALUES(UserTable.enabled.fullName, Boolean.toString(user.isEnabled()))
            .COMPOSE(MSB.when(user.getId() != -1,
                s -> s
                    .VALUES(UserTable.id.fullName, Integer.toString(user.getId()))
                    .ON_DUPLICATE_KEY_UPDATE(UserTable.username, UserTable.password, UserTable.enabled)))
            .toString(),
        save(user.getProfile()));
  }

  public String delete(String username) {
    return String.format("DELETE %s, %s FROM %s INNER JOIN %s ON %s = %s WHERE %s = %s",
        UserTable.table.name, ProfileTable.table.name,
        UserTable.table.name, ProfileTable.table.name,
        UserTable.id.fullName, ProfileTable.id.fullName,
        UserTable.username.fullName, wrapString(username));
  }

  public String changePassword(String username, String password) {
    return MyBatisSQL.create()
        .UPDATE(UserTable.table)
        .SET(UserTable.password, wrapString(password))
        .WHERE(equal(UserTable.username.fullName, wrapString(username)))
        .toString();
  }
}
