package xdean.mini.boardgame.server.endpoint;

import javax.inject.Inject;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import xdean.mini.boardgame.server.model.UserProfile;
import xdean.mini.boardgame.server.model.param.SimpleResponse;
import xdean.mini.boardgame.server.model.param.UserProfileUpdateRequest;
import xdean.mini.boardgame.server.model.param.UserProfileUpdateResponse;
import xdean.mini.boardgame.server.security.endpoint.UserAuthEndpoint;
import xdean.mini.boardgame.server.security.model.LoginResponse;

@Api(tags = "WeChat")
@RestController
@RequestMapping("/wechat")
public class WechatEndPoint {
  @Inject
  UserAuthEndpoint authEndPoint;

  @Inject
  UserProfileEndPoint profileEndPoint;

  @GetMapping("login-profile")
  @ApiOperation("login and update profile")
  public SimpleResponse loginAndUpdateProfile(
      HttpServletRequest request,
      HttpServletResponse response,
      @RequestParam("token") String token,
      @RequestParam("provider") String provider,
      @RequestParam("nickname") String nickname,
      @RequestParam("avatarUrl") String avatarUrl) {
    LoginResponse loginResponse = authEndPoint.loginOpenId(request, response, token, provider);
    if (loginResponse.getErrorCode() == 0) {
      UserProfileUpdateResponse profileResponse = profileEndPoint.updateUserProfile(UserProfileUpdateRequest.builder()
          .profile(UserProfile.builder()
              .nickname(nickname)
              .avatarUrl(avatarUrl)
              .build())
          .build());
      return SimpleResponse.builder()
          .errorCode(profileResponse.getErrorCode())
          .errorMessage("Fail to update profile")
          .build();
    } else {
      return SimpleResponse.builder()
          .errorCode(loginResponse.getErrorCode())
          .errorMessage(loginResponse.getMessage())
          .build();
    }
  }
}
