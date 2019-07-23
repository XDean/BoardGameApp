package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/wechat/model"
	"github.com/xdean/miniboardgame/go/wechat/state"
	"net/http"
	"sync"
)

var userLock = sync.Map{}
var userState = map[string]state.State{}

func Message(c echo.Context) error {
	param := new(model.Message)
	xecho.MustBindAndValidate(c, param)

	actual, _ := userLock.LoadOrStore(param.FromUserName, sync.Mutex{})
	lock := actual.(sync.Mutex)
	lock.Lock()
	s, ok := userState[param.FromUserName]
	if !ok {
		userState[param.FromUserName] = state.Root
		s = state.Root
	}
	next, msg := s.Handle(param.MsgType)(*param)
	userState[param.FromUserName] = next
	return c.XML(http.StatusOK, msg)
}
