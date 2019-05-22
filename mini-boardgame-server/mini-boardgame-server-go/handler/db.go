package handler

import (
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func DatabaseContextMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			context.Set(_const.DATABASE, db)
			return next(context)
		}
	}
}

func GetDB(e echo.Context) *gorm.DB {
	if db, ok := e.Get(_const.DATABASE).(*gorm.DB); ok {
		return db
	} else {
		panic("No db instance in context")
	}
}
