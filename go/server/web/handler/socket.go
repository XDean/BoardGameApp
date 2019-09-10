package handler

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
)

var (
	upgrader = websocket.Upgrader{}
)

func RoomSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	xecho.MustNoError(err)
	defer ws.Close()

	err = ws.WriteMessage(websocket.PingMessage, []byte("ping"))
	xecho.MustNoError(err)

	receiveStream := make(chan interface{}, 5)

	for {
		msg := <-receiveStream
		bytes, err := json.Marshal(msg)
		xecho.MustNoError(err)
		err = ws.WriteMessage(websocket.TextMessage, bytes)
		xecho.MustNoError(err)
	}
}
