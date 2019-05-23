package openid

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type OpenIdProvider struct {
	Name string
	Auth func(string) (string, error)
}

var Providers = map[string]OpenIdProvider{
	wechatOpenIdProvider.Name: wechatOpenIdProvider,
}

func Get(provider, token string) (string, error) {
	if provider, ok := Providers[provider]; ok {
		return provider.Auth(token)
	} else {
		return "", echo.NewHTTPError(http.StatusBadRequest, "No such openid provider")
	}
}
