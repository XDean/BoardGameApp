package middleware

import (
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Authorized() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if _, ok := context.Get(_const.USER).(model.User); ok {
				return next(context)
			} else {
				return &echo.HTTPError{
					Code:    http.StatusUnauthorized,
					Message: "You are not authorized",
				}
			}
		}
	}
}

func AuthRole(targetRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if user, ok := context.Get(_const.USER).(model.User); ok {
				for _, role := range user.Roles {
					if role.Name == targetRole {
						return next(context)
					}
				}
			}
			return &echo.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: "You are not " + targetRole,
			}
		}
	}
}
