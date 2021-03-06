package ocr

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/xdean/goex/xconfig"
	"github.com/xdean/miniboardgame/go/wechat/config"
	"github.com/xdean/miniboardgame/go/wechat/model"
	"github.com/xdean/miniboardgame/go/wechat/state"
	"io/ioutil"
	"net/http"
	"strings"
)

var State = OCR{
	state.BaseState{
		TheName: "图像识别文字",
		TheLast: state.Root,
	},
}

type OCR struct {
	state.BaseState
}

func (OCR) Help() string {
	return `输入图片识别图中文字`
}

func (s OCR) Handle(msgType string) state.MessageHandler {
	switch msgType {
	case model.TEXT:
		return state.DefaultText(s, func(msg model.Message) (state state.State, message model.Message) {
			return s, model.NewText(s.Help())
		})
	case model.IMAGE:
		return func(msg model.Message) (state.State, model.Message) {
			request, err := http.NewRequest("POST", "https://westcentralus.api.cognitive.microsoft.com/vision/v2.0/ocr?language=zh-Hans",
				strings.NewReader(fmt.Sprintf(`{"url":"%s"}`, msg.PicUrl)))
			if err != nil {
				return s, model.NewText("服务器错误")
			}
			request.Header.Set("Content-Type", echo.MIMEApplicationJSON)
			key, _ := xconfig.Decrypt("dftlZFppmejDaFKQN5bWKbKhTwf4mC0aMLvZ7T02dDL216+tSZIgNRsgkYivpELWixGQ5URrxF0Ax2Gs", config.SecretKey)
			request.Header.Set("Ocp-Apim-Subscription-Key", string(key))
			resp, err := http.DefaultClient.Do(request)
			if err != nil {
				return s, model.NewText("解析失败")
			} else {
				switch resp.StatusCode {
				case 400:
					return s, model.NewText("图片错误")
				case 404:
					return s, model.NewText("图片不存在")
				case 500:
					return s, model.NewText("无法解析")
				case 200:
					var body map[string]interface{}
					if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
						err = json.Unmarshal(bytes, &body)
						if regions, ok := body["regions"].([]interface{}); ok {
							builder := strings.Builder{}
							for _, v := range regions {
								if lines, ok := v.(map[string]interface{})["lines"].([]interface{}); ok {
									for _, line := range lines {
										if words, ok := line.(map[string]interface{})["words"].([]interface{}); ok {
											for _, word := range words {
												builder.WriteString(word.(map[string]interface{})["text"].(string))
											}
										}
										builder.WriteString("\n")
									}
								}
								builder.WriteString("\n\n")
							}
							return s, model.NewText(builder.String())
						}
					}
				default:
					bytes, _ := ioutil.ReadAll(resp.Body)
					logrus.Debug(resp.StatusCode, string(bytes))
				}
				return s, model.NewText("远程服务器错误")
			}
		}
	default:
		return state.HelpHandler(s)
	}
}
