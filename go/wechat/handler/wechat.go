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

//public String checkSignature(@RequestParam String signature, @RequestParam String nonce,
//@RequestParam String timestamp, @RequestParam String echostr) {
//if (WeChatUtil.checkSignature(wcv.getToken(), signature, timestamp, nonce)) {
//info("sucess");
//return echostr;
//}
//error("failed");
//return "failed";
//}

func CheckSignature(c echo.Context) error {
	type Param struct {
		Signature string `query:"signature"`
		Nonce     string `query:"nonce"`
		Timestamp string `query:"timestamp"`
		Echo      string `query:"echostr"`
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
