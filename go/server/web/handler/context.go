package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	_const "github.com/xdean/miniboardgame/go/server/const"
	"github.com/xdean/miniboardgame/go/server/model"
	"net/http"
)

func GetCurrentUser(c echo.Context) (*model.User, error) {
	if user, ok := c.Get(_const.USER).(*model.User); ok {
		return user, nil
	}
	return nil, echo.NewHTTPError(http.StatusUnauthorized, "You are not authorized")
}

func GetCurrentRoom(c echo.Context) (*model.Room, error) {
	if user, ok := c.Get(_const.ROOM).(*model.Room); ok {
		return user, nil
	}
	return nil, echo.NewHTTPError(http.StatusUnauthorized, "You are not in a room")
}

func GetDB(e echo.Context) *gorm.DB {
	if db, ok := e.Get(_const.DATABASE).(*gorm.DB); ok {
		return db
	} else {
		panic("No db instance in context")
	}
}
