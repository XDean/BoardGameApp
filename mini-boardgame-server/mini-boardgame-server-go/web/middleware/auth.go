package middleware

import (
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Authorized() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if _, ok := context.Get(_const.USER).(*model.User); ok {
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
			if user, ok := context.Get(_const.USER).(*model.User); ok {
				for _, role := range user.Roles {
					if role.Name == targetRole {
						return next(context)
					}
				}
				return &echo.HTTPError{
					Code:    http.StatusForbidden,
					Message: "You are not " + targetRole,
				}
			} else {
				return &echo.HTTPError{
					Code:    http.StatusUnauthorized,
					Message: "You are not authorized",
				}
			}
		}
	}
}

func AuthRoom() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if db, ok := context.Get(_const.DATABASE).(*gorm.DB); ok {
				if user, ok := context.Get(_const.USER).(*model.User); ok {
					room := new(model.Room)
					err := room.FindByUserID(db, user.ID)
					if err != nil {
						return echo.NewHTTPError(http.StatusBadRequest, "You are not in a room")
					}
					return next(context)
				} else {
					return echo.NewHTTPError(http.StatusUnauthorized, "You are not authorized")
				}
			} else {
				panic("There is no database in the echo context")
			}
		}
	}
}
