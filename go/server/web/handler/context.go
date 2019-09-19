package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	_const "github.com/xdean/miniboardgame/go/server/const"
	"github.com/xdean/miniboardgame/go/server/game"
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
	return nil, echo.NewHTTPError(http.StatusBadRequest, "You are not in a room")
}

func GetCurrentPlayer(c echo.Context) (*model.Player, error) {
	if player, ok := c.Get(_const.PLAYER).(*model.Player); ok {
		return player, nil
	} else if user, err := GetCurrentUser(c); err != nil {
		return nil, err
	} else {
		player := new(model.Player)
		err := player.GetByUserID(GetDB(c), user.ID)
		if err != nil {
			c.Set(_const.PLAYER, player)
		}
		return player, err
	}
}

func GetDB(e echo.Context) *gorm.DB {
	if db, ok := e.Get(_const.DATABASE).(*gorm.DB); ok {
		return db
	} else {
		panic("No db instance in context")
	}
}

func GetCurrentGame(c echo.Context) (*game.Game, error) {
	room, err := GetCurrentRoom(c)
	if err != nil {
		return nil, err
	}
	g, err := game.FindGame(room.GameId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "No such game")
	}
	return g, nil
}
