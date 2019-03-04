package xdean.mini.boardgame.server.endpoint;

import java.util.Optional;

import javax.inject.Inject;

import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import xdean.mini.boardgame.server.model.UserProfile;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mini.boardgame.server.model.param.UserProfileResponse;
import xdean.mini.boardgame.server.model.param.UserProfileUpdateRequest;
import xdean.mini.boardgame.server.model.param.UserProfileUpdateResponse;
import xdean.mini.boardgame.server.service.UserProfileRepo;
import xdean.mini.boardgame.server.service.UserService;

@RestController
@Api(tags = "User/Profile")
public class UserProfileEndPoint {

  private @Inject UserProfileRepo userProfileRepo;
  private @Inject UserService userService;

  @ApiOperation("Get user profile")
  @GetMapping(path = "/user/profile")
  public UserProfileResponse getUserProfile(@RequestParam(name = "username", required = false) String username) {
    if (username == null) {
      username = userService.getCurrentUser().map(u -> u.getUsername()).orElse(null);
    }
    if (username == null) {
      return UserProfileResponse.builder().errorCode(UserProfileResponse.INPUT_USER).build();
    }
    Optional<UserProfileEntity> p = userProfileRepo.findByUserUsername(username);
    if (p.isPresent()) {
      UserProfileEntity profile = p.get();
      if (profile == null) {
        return UserProfileResponse.builder().errorCode(UserProfileResponse.PROFILE_NOT_FOUND).build();
      }
      return UserProfileResponse.builder().profile(profile.getProfile()).build();
    } else {
      return UserProfileResponse.builder().errorCode(UserProfileResponse.USER_NOT_FOUND).build();
    }
  }

  @ApiOperation("Update user profile")
  @PostMapping(path = "/user/profile")
  public UserProfileUpdateResponse updateUserProfile(@Validated @RequestBody UserProfileUpdateRequest request) {
    String username = userService.getCurrentUser().map(u -> u.getUsername()).orElse(null);
    if (username == null) {
      return UserProfileUpdateResponse.builder().errorCode(UserProfileUpdateResponse.HAVE_NOT_LOGIN).build();
    }
    Optional<UserProfileEntity> u = userProfileRepo.findByUserUsername(username);
    if (!u.isPresent()) {
      throw new IllegalStateException("A user loged in but there is no record in DB");
    }
    UserProfileEntity p = userProfileRepo.save(u.get().toBuilder()
        .profile(UserProfile.builder()
            .male(request.getProfile().isMale())
            .nickname(request.getProfile().getNickname())
            .avatarUrl(request.getProfile().getAvatarUrl())
            .build())
        .build());
    return UserProfileUpdateResponse.builder().profile(p.getProfile()).build();
  }
}
