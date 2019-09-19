package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/game"
	"net/http"
)

func GetGameList(c echo.Context) error {
	return c.JSON(http.StatusOK, gamesJson(game.Games))
}

func StartGame(c echo.Context) error {
	user, err := GetCurrentUser(c)
	xecho.MustNoError(err)

	room, err := GetCurrentRoom(c)
	xecho.MustNoError(err)

	g, err := GetCurrentGame(c)
	xecho.MustNoError(err)

	if room.FindHost().UserID != user.ID {
		return echo.NewHTTPError(http.StatusBadRequest, xecho.M("You are not host"))
	}
	if !room.IsAllReady() {
		return echo.NewHTTPError(http.StatusBadRequest, xecho.M("Players not ready"))
	}
	res := g.OnEvent(game.NewGameEvent{
		BaseEvent: game.BaseEvent{
			ResponseStream: make(chan game.Response),
			User:           user,
		},
	})
	return eventResponse(c, res)
}

func GameEvent(c echo.Context) error {
	g, err := GetCurrentGame(c)
	xecho.MustNoError(err)

	e := g.NewEvent()
	xecho.MustBindAndValidate(c, &e)
	res := g.OnEvent(e)
	return eventResponse(c, res)
}

func eventResponse(c echo.Context, res game.Response) error {
	switch t := res.(type) {
	case error:
		return c.JSON(http.StatusBadRequest, xecho.M(t.Error()))
	case string:
		return c.JSON(http.StatusOK, xecho.M(t))
	default:
		return c.JSON(http.StatusOK, t)
	}
}

func gamesJson(games []*game.Game) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range games {
		result = append(result, gameJson(v))
	}
	return result
}

func gameJson(v *game.Game) interface{} {
	return xecho.J{
		"Id":       v.Id,
		"Name":     v.Name,
		"Player":   v.Player,
		"Options:": v.Options,
	}
}
