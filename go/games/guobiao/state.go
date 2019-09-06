package guobiao

import (
	"fmt"
	"github.com/xdean/miniboardgame/go/wechat/model"
	. "github.com/xdean/miniboardgame/go/wechat/state"
	"strings"
)

var StateInstance = GuoBiao{
	BaseState{
		TheName: "国标麻将算番",
		TheLast: Root,
	},
}

type GuoBiao struct {
	BaseState
}

func (s GuoBiao) Help() string {
	return `输入牌型算番
牌型表示为
> 明牌，暗牌，胡牌
例如
> 明东南 暗西北，中中，中
> 吃条123 饼234 碰万4，万34588，摸8
> 碰发 条1，饼222 万333 北北，北
特殊番型需另计：如和绝张(4)、抢杠和(8)、海底捞月(8)等`
}

func (s GuoBiao) Handle(msgType string) MessageHandler {
	switch msgType {
	case model.TEXT:
		return DefaultText(s, func(input model.Message) (state State, message model.Message) {
			hand, err := Parse(input.Content)
			if err != nil {
				return s, model.NewText(err.Error())
			}
			fan := CalcFan(hand)
			b := strings.Builder{}
			for _, v := range fan {
				b.WriteString(v.String())
				b.WriteString("\n")
			}
			b.WriteString(fmt.Sprintf("共计 %d 番", fan.Sum()))
			return s, model.NewText(b.String())
		})
	default:
		return HelpHandler(s)
	}
}
