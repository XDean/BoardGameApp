package handler

import (
	"github.com/XDean/MiniBoardgame/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateRoom(c echo.Context) error {
	type Param struct {
		GameName    string `json:"game_name" query:"game_name" form:"game_name" validate:"required"`
		RoomName    string `json:"room_name" query:"room_name" form:"room_name" validate:"required"`
		PlayerCount uint   `json:"player_count" query:"player_count" form:"player_count" validate:"required"`
	}
	param := new(Param)
	BindAndValidate(c, param)

	user, err := GetCurrentUser(c)
	MustNoError(err)

	db := GetDB(c)

	player := new(model.Player)
	err = player.GetByUserID(db, user.ID)
	MustNoError(err)

	if player.RoomID != 0 {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "You have been in a room")
	}

	room := new(model.Room)
	room.GameName = param.GameName
	room.RoomName = param.RoomName
	room.PlayerCount = param.PlayerCount
	err = room.CreateByHost(db, user)
	MustNoError(err)

	return c.JSON(http.StatusOK, room)
}

func GetRoom(c echo.Context) error {
	user, err := GetCurrentUser(c)
	MustNoError(err)

	room := new(model.Room)
	err = room.FindByUserID(GetDB(c), user.ID)
	if err != nil {
		return DBNotFound(err, "You are not in a room")
	}
	return c.JSON(http.StatusOK, room)
}
