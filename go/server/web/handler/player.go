package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/model"
	"net/http"
	"strconv"
)

func GetPlayer(c echo.Context) error {
	user, err := GetCurrentUser(c)
	xecho.MustNoError(err)

	player := new(model.Player)
	err = player.GetByUserID(GetDB(c), user.ID)
	xecho.MustNoError(err)

	return c.JSON(http.StatusOK, playerJson(player))
}

func GetPlayerByID(c echo.Context) error {
	idParam := c.Param("id")
	if id, err := strconv.Atoi(idParam); err == nil {
		player := new(model.Player)
		err = player.GetByUserID(GetDB(c), uint(id))
		xecho.MustNoError(err)
		return c.JSON(http.StatusOK, playerJson(player))
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Unrecognized id: "+idParam)
	}
}

func playerJson(p *model.Player) interface{} {
	return xecho.J{
		"UserID":      p.UserID,
		"State":       p.State,
		"StateString": p.State.String(),
		"RoomId":      p.RoomID,
		"Seat":        p.Seat,
	}
}
