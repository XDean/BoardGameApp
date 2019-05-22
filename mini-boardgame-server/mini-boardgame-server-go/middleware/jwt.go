package middleware

import (
	"fmt"
	"github.com/XDean/MiniBoardgame/config"
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	jwtExtractor func(echo.Context) (string, bool)
)

const (
	AlgorithmHS256 = "HS256"
	JwtKey         = "access-token"
)

func Jwt() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		key := []byte(config.Global.Security.Key)
		return func(c echo.Context) error {
			var auth string
			var ok bool
			var err error
			auth, ok = jwtFromQuery(JwtKey)(c)
			if !ok {
				auth, ok = jwtFromHeader(echo.HeaderAuthorization)(c)
			}
			if !ok {
				auth, ok = jwtFromCookie(JwtKey)(c)
			}
			if !ok {
				return next(c)
			}
			token := new(jwt.Token)
			token, err = jwt.ParseWithClaims(auth, &model.Claims{}, func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != AlgorithmHS256 {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return key, nil
			})
			if err == nil && token.Valid {
				claims := token.Claims.(*model.Claims)
				user := model.User{
					ID:       claims.UserID,
					Username: claims.Username,
				}
				user.SetRoles(claims.Roles)
				c.Set(_const.USER, user)
			}
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "invalid or expired jwt",
				Internal: err,
			}
		}
	}
}

func jwtFromHeader(header string) jwtExtractor {
	return func(c echo.Context) (string, bool) {
		if auth := c.Request().Header.Get(header); auth == "" {
			return auth, false
		} else {
			return auth, true
		}
	}
}

func jwtFromQuery(param string) jwtExtractor {
	return func(c echo.Context) (string, bool) {
		token := c.QueryParam(param)
		if token == "" {
			return "", false
		}
		return token, true
	}
}

func jwtFromCookie(name string) jwtExtractor {
	return func(c echo.Context) (string, bool) {
		cookie, err := c.Cookie(name)
		if err != nil {
			return "", false
		}
		return cookie.Value, true
	}
}
