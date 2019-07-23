package state

import (
	"fmt"
	"github.com/xdean/miniboardgame/go/wechat/model"
)

type (
	MessageHandler func(input model.Message) (State, model.Message)
	State          interface {
		Name() string
		String() string // name
		Help() string
		Handle(msgType string) MessageHandler
	}

	BaseState struct {
		name string
		last State
	}

	RootState struct {
		BaseState
	}
)

var Root = RootState{BaseState{name: "root", last: nil}}
var rootStates = make([]State, 0)
var rootList = OptionList{
	Options: []fmt.Stringer{},
}

func Register(state State) {
	rootStates = append(rootStates, state)
	rootList.Options = append(rootList.Options, state)
}

func (s BaseState) Name() string {
	return s.name
}

func (s BaseState) String() string {
	return s.name
}

func (s RootState) Help() string {
	return rootList.String()
}

func (s RootState) Handle(msgType string) MessageHandler {
	return handleOrHelp(s, model.TEXT, msgType, func(input model.Message) (state State, message model.Message) {
		c := rootList.Find(input.Content)
		if c == nil {
			return s, model.NewText(s.Help())
		} else {
			next := c.(State)
			return next, model.NewText(next.Help())
		}
	})
}

func handleOrHelp(s State, targetType, msgType string, handler MessageHandler) MessageHandler {
	if targetType == msgType {
		return handler
	} else {
		return func(input model.Message) (state State, message model.Message) {
			return s, model.NewText(s.Help())
		}
	}
}
