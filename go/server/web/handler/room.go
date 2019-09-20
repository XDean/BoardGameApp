package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/model"
	"net/http"
	"strconv"
)

func CreateRoom(c echo.Context) error {
	type Param struct {
		GameId      string `json:"game_id" query:"game_id" form:"game_id" validate:"required"`
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
	room.GameId = param.GameId
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

func GetRoomByID(c echo.Context) error {
	room := findRoomByID(c)
	return c.JSON(http.StatusOK, roomJson(room))
}

func JoinRoom(c echo.Context) error {
	player, err := GetCurrentPlayer(c)
	xecho.MustNoError(err)

	if player.IsInGame() {
		return c.JSON(http.StatusBadRequest, xecho.M("You are in room. Exit first."))
	}

	room := findRoomByID(c)

	resultStream := make(chan error)

	room.Do(func() {
		if len(room.Players) < int(room.PlayerCount) {
			if err := room.Join(GetDB(c), player); err != nil {
				resultStream <- err
			} else {
				resultStream <- c.JSON(http.StatusOK, xecho.M("Join success"))
			}
		} else {
			resultStream <- echo.NewHTTPError(http.StatusBadRequest, "The room is full")
		}
	})

	return <-resultStream
}

func ExitRoom(c echo.Context) error {
	player, err := GetCurrentPlayer(c)
	xecho.MustNoError(err)

	if !player.IsInGame() {
		return c.JSON(http.StatusBadRequest, xecho.M("You are not in room"))
	}

	resultStream := make(chan error)
	player.Room.Do(func() {
		err = player.Room.Exit(GetDB(c), player)
		if err != nil {
			resultStream <- err
		} else {
			resultStream <- c.JSON(http.StatusOK, xecho.M("Exit success"))
		}
	})

	return <-resultStream
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
		//room.SendEvent()
		// TODO  there is no player, go to the seat directly
	}
	return nil
}

func roomJson(room *model.Room) interface{} {
	players := make([]interface{}, 0)
	for _, player := range room.Players {
		players = append(players, roomPlayerJson(player))
	}
	return xecho.J{
		"ID":          room.ID,
		"GameId":      room.GameId,
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

func findRoomByID(c echo.Context) *model.Room {
	idParam := c.Param("id")
	if id, err := strconv.Atoi(idParam); err == nil {
		room := new(model.Room)
		if err = room.FindByID(GetDB(c), uint(id)); err != nil {
			xecho.MustNoError(DBNotFound(err, "No such room"))
		}
		return room
	} else {
		xecho.MustNoError(echo.NewHTTPError(http.StatusBadRequest, "Unrecognized id: "+idParam))
	}
	return nil
}