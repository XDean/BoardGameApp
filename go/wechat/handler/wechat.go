package handler

import (
	"crypto/sha1"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/wechat/config"
	"net/http"
	"sort"
	"strings"
)

func CheckSignature(c echo.Context) error {
	type Param struct {
		Signature string `query:"signature" validate:"required"`
		Nonce     string `query:"nonce" validate:"required"`
		Timestamp string `query:"timestamp" validate:"required"`
		Echo      string `query:"echostr" validate:"required"`
	}

	param := new(Param)
	xecho.MustBindAndValidate(c, param)

	if checkSignature(config.Instance.Wechat.Token, param.Signature, param.Timestamp, param.Nonce) {
		return c.String(http.StatusOK, param.Echo)
	} else {
		return c.String(http.StatusBadRequest, "Bad Signature")
	}
}

func checkSignature(token, signature, timestamp, nonce string) bool {
	array := []string{token, timestamp, nonce}
	sort.Strings(array)
	join := strings.Join(array, "")
	sum := sha1.Sum([]byte(join))
	return signature == fmt.Sprintf("%x", sum)
}
