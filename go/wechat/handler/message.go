package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/games/guobiao/guobiao"
	"github.com/xdean/miniboardgame/go/wechat/model"
	"net/http"
	"time"
)

func Message(c echo.Context) error {
	param := new(model.Message)
	xecho.MustBindAndValidate(c, param)

	hand, err := guobiao.Parse(param.Content)
	if err != nil {
		return err
	}
	fan := guobiao.CalcFan(hand)

	return c.XML(http.StatusOK, model.Message{
		FromUserName: param.ToUserName,
		ToUserName:   param.FromUserName,
		CreateTime:   time.Now().Unix(),
		Content:      fmt.Sprintf("番型: %s", fan),
		MsgType:      model.TEXT,
	})
}
