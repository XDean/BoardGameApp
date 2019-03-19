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
import xdean.mini.boardgame.server.model.GlobalConstants.AttrKey;
import xdean.mini.boardgame.server.model.entity.UserProfileEntity;
import xdean.mini.boardgame.server.model.param.UserProfileUpdateRequest;
import xdean.mini.boardgame.server.model.param.UserProfileUpdateResponse;
import xdean.mini.boardgame.server.security.endpoint.UserAuthEndpoint;

@Api(tags = "WeChat")
@RestController
@RequestMapping("/public/wechat")
public class WechatEndPoint {
  @Inject
  UserAuthEndpoint authEndPoint;

  @Inject
  UserProfileEndPoint profileEndPoint;

  @GetMapping("login-profile")
  @ApiOperation("login and update profile")
  public UserProfileUpdateResponse loginAndUpdateProfile(
      HttpServletRequest request,
      HttpServletResponse response,
      @RequestParam("token") String token,
      @RequestParam("provider") String provider,
      @RequestParam("nickname") String nickname,
      @RequestParam("avatarUrl") String avatarUrl) {
    authEndPoint.loginOpenId(request, response, token, provider);
    Integer userId = (Integer) request.getSession().getAttribute(AttrKey.USER_ID);
    return profileEndPoint.updateUserProfile(userId, UserProfileUpdateRequest.builder()
        .profile(UserProfileEntity.builder()
            .nickname(nickname)
            .avatarUrl(avatarUrl)
            .build())
        .build());
  }
}
