package state

type Command string

var (
	Back Command = "返回"
	Help Command = "帮助"
)

func (c Command) String() string {
	return string(c)
}
