package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/model"
	"net/http"
)

func CreateRoom(c echo.Context) error {
	type Param struct {
		GameName    string `json:"game_name" query:"game_name" form:"game_name" validate:"required"`
		RoomName    string `json:"room_name" query:"room_name" form:"room_name" validate:"required"`
		PlayerCount uint   `json:"player_count" query:"player_count" form:"player_count" validate:"required"`
	}
	param := new(Param)
	xecho.MustBindAndValidate(c, param)

	user, err := GetCurrentUser(c)
	xecho.MustNoError(err)

	db := GetDB(c)

	player := new(model.Player)
	err = player.GetByUserID(db, user.ID)
	xecho.MustNoError(err)

	if player.RoomID != 0 {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "You have been in a room")
	}

	room := new(model.Room)
	room.GameName = param.GameName
	room.RoomName = param.RoomName
	room.PlayerCount = param.PlayerCount
	err = room.CreateByHost(db, player)
	xecho.MustNoError(err)

	return c.JSON(http.StatusOK, roomJson(room))
}

func GetRoom(c echo.Context) error {
	room, err := GetCurrentRoom(c)
	xecho.MustNoError(err)
	return c.JSON(http.StatusOK, roomJson(room))
}

func roomJson(room *model.Room) interface{} {
	players := make([]interface{}, 0)
	for _, player := range room.Players {
		players = append(players, roomPlayerJson(player))
	}
	return xecho.J{
		"ID":          room.ID,
		"GameName":    room.GameName,
		"RoomName":    room.RoomName,
		"PlayerCount": room.PlayerCount,
		"CreatedTime": room.CreatedTime,
		"Players":     players,
	}
}

func roomPlayerJson(player *model.Player) interface{} {
	return xecho.J{
		"UserID":      player.UserID,
		"State":       player.State,
		"StateString": player.State.String(),
		"Seat":        player.Seat,
	}
}

func SwapSeat(c echo.Context) error {
	user, err := GetCurrentUser(c)
	xecho.MustNoError(err)

	room, err := GetCurrentRoom(c)
	xecho.MustNoError(err)

	targetSeat := IntParam(c, "seat")
	if player, ok := room.FindPlayerBySeat(uint(targetSeat)); ok {
		if player.UserID == user.ID {
			return echo.NewHTTPError(http.StatusBadRequest, "Can't swap seat with yourself")
		}
		//TODO
		// go swap()
	} else {
		// TODO  there is no player, go to the seat directly
	}
	return nil
}
