package middleware

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	_const "github.com/xdean/miniboardgame/go/server/const"
)

func DbContext(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			context.Set(_const.DATABASE, db)
			return next(context)
		}
	}
}
