package xdean.mini.boardgame.server.mybatis.mapper;

import static xdean.mybatis.extension.SqlUtil.wrapString;

import java.util.List;

import org.apache.ibatis.annotations.Mapper;

import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mini.boardgame.server.mybatis.Tables;
import xdean.mybatis.extension.annotation.DefaultBuilder;
import xdean.mybatis.extension.annotation.ResultMapType;

@Mapper
@DefaultBuilder(UserMapperBuilder.class)
public interface UserMapper extends BaseMapper, Tables {

  @ResultMapType
  UserEntity findByUsername(String username);

  @ResultMapType
  UserEntity findUserById(Integer id);

  List<String> findAuthorities(Integer id);

  void saveAuthorities(UserEntity user);

  void saveProfile(UserProfileEntity profile);

  int createUser(UserEntity user);

  void updateUser(UserEntity user);

  void delete(String username);

  void changePassword(String username, String password);

  default boolean userExist(int id) {
    return exist(UserTable.id, Integer.toString(id));
  }

  default boolean userExist(String username) {
    return exist(UserTable.username, wrapString(username));
  }
}
