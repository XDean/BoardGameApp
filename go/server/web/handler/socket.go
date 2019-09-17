package handler

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	topic "github.com/xdean/miniboardgame/go/server/const/socket"
	"github.com/xdean/miniboardgame/go/server/model/space"
)

var (
	upgrader = websocket.Upgrader{}
)

func RoomSocket(c echo.Context) error {
	user, err := GetCurrentUser(c)
	xecho.MustNoError(err)

	room, err := GetCurrentRoom(c)
	xecho.MustNoError(err)

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	xecho.MustNoError(err)
	defer ws.Close()

	err = ws.WriteMessage(websocket.PingMessage, []byte("ping"))
	xecho.MustNoError(err)

	room.SendEvent(space.Event{
		From:    int(user.ID),
		To:      -1,
		Topic:   topic.PLAYER_CONNECTED,
		Payload: user.ID,
	})

	subscription := room.Listen()

	ws.SetCloseHandler(func(code int, text string) error {
		room.SendEvent(space.Event{
			From:    int(user.ID),
			To:      -1,
			Topic:   topic.PLAYER_DISCONNECTED,
			Payload: user.ID,
		})
		subscription.Done <- true
		return nil
	})

	for {
		e, ok := <-subscription.EventListener
		if !ok {
			return nil
		}
		bytes, err := json.Marshal(e)
		xecho.MustNoError(err)

		err = ws.WriteMessage(websocket.TextMessage, bytes)
		xecho.MustNoError(err)
	}
}
