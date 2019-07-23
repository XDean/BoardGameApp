package state

import "fmt"

type (
	State interface {
		String() string
		Handle(input string) (State, string)
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

func Register(state State) {
	rootStates = append(rootStates, state)
}

func (s BaseState) String() string {
	return s.name
}

func (s RootState) Handle(input string) (State, string) {
	list := OptionList{
		Options: []fmt.Stringer{},
	}
	if input == "" {
		return s, list.String()
	}
	for _, v := range rootStates {
		list.Options = append(list.Options, v)
	}
	c := list.Find(input)
	if c == nil {
		return s, list.String()
	} else {
		next := c.(State)
		_, text := next.Handle("")
		return next, text
	}
}
