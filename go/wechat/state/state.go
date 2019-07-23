package state

import (
	"fmt"
	"github.com/xdean/miniboardgame/go/wechat/model"
)

type (
	MessageHandler func(msg model.Message) (State, model.Message)
	State          interface {
		Name() string
		String() string // name
		Help() string
		Last() State
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

func (s BaseState) Last() State {
	return s.last
}

func (s RootState) Help() string {
	return rootList.String()
}

func (s RootState) Handle(msgType string) MessageHandler {
	switch msgType {
	case model.TEXT:
		return defaultText(s, func(input model.Message) (state State, message model.Message) {
			c := rootList.Find(input.Content)
			if c == nil {
				return s, model.NewText(s.Help())
			} else {
				next := c.(State)
				return next, model.NewText(next.Help())
			}
		})
	default:
		return helpHandler(s)
	}
}

func defaultText(s State, h MessageHandler) MessageHandler {
	return func(input model.Message) (state State, message model.Message) {
		switch input.Content {
		case Help.String():
			return s, model.NewText(s.Help())
		case Back.String():
			last := s.Last()
			if last != nil {
				return last, model.NewText(last.Help())
			}
		}
		return h(input)
	}
}

func helpHandler(s State) MessageHandler {
	return func(input model.Message) (state State, message model.Message) {
		return s, model.NewText(s.Help())
	}
}
