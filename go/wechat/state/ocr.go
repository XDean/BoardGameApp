package state

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/xdean/miniboardgame/go/wechat/model"
	"io/ioutil"
	"net/http"
	"strings"
)

func init() {
	Register(OCR{
		BaseState{
			name: "图像识别文字",
			last: Root,
		},
	})
}

type OCR struct {
	BaseState
}

func (OCR) Help() string {
	return `输入图片识别图中文字`
}

func (s OCR) Handle(msgType string) MessageHandler {
	switch msgType {
	case model.IMAGE:
		return func(msg model.Message) (State, model.Message) {
			resp, err := http.Post("https://westcentralus.api.cognitive.microsoft.com/vision/v2.0/ocr?language=zh-Hans", echo.MIMEApplicationJSON,
				strings.NewReader(fmt.Sprintf(`{"url":"%s"}`, msg.PicUrl)))
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
						if regions, ok := body["regions"].([]map[string]interface{}); ok {
							builder := strings.Builder{}
							for _, v := range regions {
								if lines, ok := v["lines"].([]map[string]interface{}); ok {
									for _, line := range lines {
										if words, ok := line["words"].([]map[string]interface{}); ok {
											for _, word := range words {
												builder.WriteString(word["text"].(string))
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
				}
				return s, model.NewText("远程服务器错误")
			}
		}
	default:
		return helpHandler(s)
	}
}
