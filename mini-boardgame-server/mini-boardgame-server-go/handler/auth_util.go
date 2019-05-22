package handler

import (
	"errors"
	"github.com/XDean/MiniBoardgame/config"
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func GetCurrentUser(c echo.Context) (*model.User, error) {
	if user, ok := c.Get(_const.USER).(*model.User); ok {
		return user, nil
	}
	return nil, errors.New("not authorized")
}

type Claims struct {
	jwt.StandardClaims
	User model.User
}

func AuthMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.Global.Security.Key),
		Claims:     Claims{},
		SuccessHandler: func(echo echo.Context) {
			token := echo.Get(middleware.DefaultJWTConfig.ContextKey).(jwt.Token)
			claims := token.Claims.(Claims)
			echo.Set(_const.USER, claims.User)
		},
	})
}

func AdminAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		if user, ok := context.Get(_const.USER).(model.User); ok {
			for _, role := range user.Roles {
				if _const.ROLE_ADMIN == role.Name {
					return next(context)
				}
			}
		}
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "not authorized or not admin",
		}
	}
}
