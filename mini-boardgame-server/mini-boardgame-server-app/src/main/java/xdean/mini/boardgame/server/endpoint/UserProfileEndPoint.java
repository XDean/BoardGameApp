package xdean.mini.boardgame.server.endpoint;

import java.util.Optional;

import javax.inject.Inject;

import org.springframework.http.HttpStatus;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.SessionAttribute;

import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import springfox.documentation.annotations.ApiIgnore;
import xdean.mini.boardgame.server.model.CommonConstants.AttrKey;
import xdean.mini.boardgame.server.model.entity.UserEntity;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mini.boardgame.server.model.exception.MiniBoardgameException;
import xdean.mini.boardgame.server.model.param.UserProfileResponse;
import xdean.mini.boardgame.server.model.param.UserProfileUpdateRequest;
import xdean.mini.boardgame.server.model.param.UserProfileUpdateResponse;
import xdean.mini.boardgame.server.service.UserDataService;

@RestController
@Api(tags = "User/Profile")
public class UserProfileEndPoint {

  private @Inject UserDataService userService;

  @ApiOperation("Get user profile")
  @GetMapping(path = "/user/profile")
  public UserProfileResponse getUserProfile(@RequestParam(name = "id", required = false) Integer id) {
    Optional<UserEntity> user;
    if (id == null) {
      user = userService.getCurrentUser();
    } else {
      user = userService.findUserById(id);
    }
    if (user.isPresent()) {
      UserProfileEntity profile = user.get().getProfile();
      if (profile == null) {
        profile = UserProfileEntity.builder().build();
      }
      return UserProfileResponse.builder()
          .profile(profile)
          .build();
    } else {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.NOT_FOUND)
          .message("No such user")
          .build();
    }
  }

  @ApiOperation("Update user profile")
  @PostMapping(path = "/user/profile")
  public UserProfileUpdateResponse updateUserProfile(@ApiIgnore @SessionAttribute(AttrKey.USER_ID) int userId,
      @Validated @RequestBody UserProfileUpdateRequest request) {
    String username = userService.getCurrentUser().map(u -> u.getUsername()).orElse(null);
    if (username == null) {
      throw MiniBoardgameException.builder()
          .code(HttpStatus.UNAUTHORIZED)
          .message("No authorized user")
          .build();
    }
    UserProfileEntity p = request.getProfile().toBuilder().id(userId).build();
    userService.save(p);
    return UserProfileUpdateResponse.builder().profile(p).build();
  }
}
