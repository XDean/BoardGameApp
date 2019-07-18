package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/wechat/model"
	"net/http"
	"time"
)

func Message(c echo.Context) error {
	param := new(model.Message)
	xecho.MustBindAndValidate(c, param)

	return c.XML(http.StatusOK, model.Message{
		FromUserName: param.ToUserName,
		ToUserName:   param.FromUserName,
		CreateTime:   time.Now().Unix(),
		Content:      "echo: " + param.Content,
		MsgType:      model.TEXT,
	})
}
