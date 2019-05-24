package handler

import (
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCurrentUser(c echo.Context) (*model.User, error) {
	if user, ok := c.Get(_const.USER).(*model.User); ok {
		return user, nil
	}
	return nil, echo.NewHTTPError(http.StatusUnauthorized, "You are not authorized")
}

func GetDB(e echo.Context) *gorm.DB {
	if db, ok := e.Get(_const.DATABASE).(*gorm.DB); ok {
		return db
	} else {
		panic("No db instance in context")
	}
}
