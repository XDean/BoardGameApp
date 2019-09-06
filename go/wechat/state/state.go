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
		TheName string
		TheLast State
	}

	RootState struct {
		BaseState
	}
)

var Root = RootState{BaseState{TheName: "root", TheLast: nil}}
var rootStates = make([]State, 0)
var rootList = OptionList{
	Options: []fmt.Stringer{},
}

func Register(state State) {
	if state.Last() != Root {
		panic("RootState's Last must be Root")
	}
	rootStates = append(rootStates, state)
	rootList.Options = append(rootList.Options, state)
}

func (s BaseState) Name() string {
	return s.TheName
}

func (s BaseState) String() string {
	return s.TheName
}

func (s BaseState) Last() State {
	return s.TheLast
}

func (s RootState) Help() string {
	return rootList.String()
}

func (s RootState) Handle(msgType string) MessageHandler {
	switch msgType {
	case model.TEXT:
		return DefaultText(s, func(input model.Message) (state State, message model.Message) {
			c := rootList.Find(input.Content)
			if c == nil {
				return s, model.NewText(s.Help())
			} else {
				next := c.(State)
				return next, model.NewText(next.Help())
			}
		})
	default:
		return HelpHandler(s)
	}
}

func DefaultText(s State, h MessageHandler) MessageHandler {
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

func HelpHandler(s State) MessageHandler {
	return func(input model.Message) (state State, message model.Message) {
		return s, model.NewText(s.Help())
	}
}
