package xdean.mini.boardgame.server.model;

import org.apache.ibatis.type.JdbcType;

import xdean.mybatis.extension.model.Column;
import xdean.mybatis.extension.model.Table;
import xdean.mybatis.extension.model.Column.ColumnBuilder;

public interface Tables {
  interface UserTable {
    Table table = Table.create("t_users");

    Column id = ColumnBuilder.create().table(table).column("id").jdbcType(JdbcType.INTEGER).id().build();
    Column username = ColumnBuilder.create().table(table).column("username").jdbcType(JdbcType.VARCHAR).build();
    Column password = ColumnBuilder.create().table(table).column("password").jdbcType(JdbcType.VARCHAR).build();
    Column enabled = ColumnBuilder.create().table(table).column("enabled").jdbcType(JdbcType.BOOLEAN).build();
  }

  interface AuthorityTable {
    Table table = Table.create("t_authorities");

    Column id = ColumnBuilder.create().table(table).column("id").jdbcType(JdbcType.INTEGER).build();
    Column authority = ColumnBuilder.create().table(table).column("authority").jdbcType(JdbcType.VARCHAR).build();
  }

  interface ProfileTable {
    Table table = Table.create("t_user_profile");

    Column id = ColumnBuilder.create().table(table).column("id").jdbcType(JdbcType.INTEGER).id().build();
    Column avatarUrl = ColumnBuilder.create().table(table).column("avatar_url").jdbcType(JdbcType.VARCHAR).build();
    Column male = ColumnBuilder.create().table(table).column("male").jdbcType(JdbcType.BOOLEAN).build();
    Column nickname = ColumnBuilder.create().table(table).column("nickname").jdbcType(JdbcType.VARCHAR).build();
  }

  interface GamePlayerTable {
    Table table = Table.create("t_game_players");

    Column id = ColumnBuilder.create().table(table).column("id").jdbcType(JdbcType.INTEGER).id().build();
    Column roomId = ColumnBuilder.create().table(table).column("room_id").jdbcType(JdbcType.INTEGER).index().build();
    Column seat = ColumnBuilder.create().table(table).column("seat").jdbcType(JdbcType.INTEGER).build();
  }

  interface GameRoomTable {
    Table table = Table.create("t_game_rooms");

    Column id = ColumnBuilder.create().table(table).column("id").jdbcType(JdbcType.INTEGER).id().build();
    Column gameName = ColumnBuilder.create().table(table).column("game_name").jdbcType(JdbcType.VARCHAR).index().build();
    Column roomName = ColumnBuilder.create().table(table).column("room_name").jdbcType(JdbcType.VARCHAR).build();
    Column playerCount = ColumnBuilder.create().table(table).column("player_count").jdbcType(JdbcType.INTEGER).build();
    Column createdTime = ColumnBuilder.create().table(table).column("board_time").jdbcType(JdbcType.TIMESTAMP).build();
    Column board = ColumnBuilder.create().table(table).column("board").jdbcType(JdbcType.VARCHAR).build();
  }
}
