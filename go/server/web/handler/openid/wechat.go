package openid

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/xdean/miniboardgame/go/server/config"
	"gopkg.in/resty.v1"
	"net/http"
)

type wechatAuthInfo struct {
	OpenId       string `json:"openid"`
	SessionKey   string `json:"session_key"`
	UnionId      string `json:"unionid"`
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

var wechatOpenIdProvider = OpenIdProvider{
	Name: "wechat-mbg",
	Auth: func(token string) (string, error) {
		try := 0
	TRY:
		try++
		response, err := resty.R().SetQueryParams(map[string]string{
			"appid":      config.Global.Wechat.AppId,
			"secret":     config.Global.Wechat.AppSecret,
			"js_code":    token,
			"grant_type": "authorization_code",
		}).Get(config.Global.Wechat.AuthUrl)
		if err == nil {
			body := response.Body()
			info := new(wechatAuthInfo)
			err = json.Unmarshal(body, info)
			if err == nil {
				switch info.ErrorCode {
				case 0:
					return info.OpenId, nil
				case -1:
					if try > 5 {
						return "", echo.NewHTTPError(http.StatusServiceUnavailable, "Server busy, try again later")
					} else {
						goto TRY
					}
				case 40029:
					return "", echo.NewHTTPError(http.StatusUnauthorized, "Wrong wechat mini boardgame token")
				case 45011:
					return "", echo.NewHTTPError(http.StatusTooManyRequests, "Wechat server limited")
				default:
					return "", echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unknown error code: %d", info.ErrorCode))
				}
			} else {
				return "", err
			}
		} else {
			return "", err
		}
	},
}
