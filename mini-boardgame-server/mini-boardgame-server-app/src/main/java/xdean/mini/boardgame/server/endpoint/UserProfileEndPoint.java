package xdean.mini.boardgame.server.endpoint;

import java.util.Optional;

import javax.inject.Inject;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import xdean.mini.boardgame.server.model.UserProfile;
import xdean.mini.boardgame.server.model.param.UserProfileResponse;
import xdean.mini.boardgame.server.model.param.UserProfileUpdateResponse;
import xdean.mini.boardgame.server.service.UserProfileRepo;
import xdean.mini.boardgame.server.util.UserUtil;

@RestController
@Api(tags = "User/Profile")
public class UserProfileEndPoint {

  @Inject
  UserProfileRepo userProfileRepo;

  @ApiOperation("Get user profile")
  @GetMapping(path = "/user/profile")
  public UserProfileResponse getUserProfile(@RequestParam(name = "username", required = false) String username) {
    if (username == null) {
      username = UserUtil.getAuthUser().map(u -> u.getUsername()).orElse(null);
    }
    if (username == null) {
      return UserProfileResponse.builder().errorCode(UserProfileResponse.INPUT_USER).build();
    }
    Optional<UserProfile> p = userProfileRepo.findById(username);
    if (p.isPresent()) {
      return UserProfileResponse.builder().profile(p.get()).build();
    } else {
      return UserProfileResponse.builder().errorCode(UserProfileResponse.USER_NOT_FOUND).build();
    }
  }

  @ApiOperation("Update user profile")
  @PostMapping(path = "/user/profile-set")
  public UserProfileUpdateResponse updateUserProfile(@RequestBody UserProfile profile) {
    String username = UserUtil.getAuthUser().map(u -> u.getUsername()).orElse(null);
    if (username == null || !username.equals(profile.getUsername())) {
      return UserProfileUpdateResponse.builder().errorCode(UserProfileUpdateResponse.NOT_CURRENT_USER).build();
    }
    UserProfile p = userProfileRepo.save(profile);
    return UserProfileUpdateResponse.builder().profile(p).build();
  }
}