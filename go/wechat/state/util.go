package state

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	OptionList struct {
		What    string
		Options []fmt.Stringer
	}
)

func (o OptionList) String() string {
	b := strings.Builder{}
	b.WriteString("请选择")
	b.WriteString(o.What)
	b.WriteString("\n")
	for i, v := range o.Options {
		b.WriteString(strconv.Itoa(i) + ". ")
		b.WriteString(v.String())
		b.WriteString("\n")
	}
	b.WriteString("(输入选项或者序号)")
	return b.String()
}

func (o OptionList) Find(input string) fmt.Stringer {
	for i, v := range o.Options {
		if input == strconv.Itoa(i) || input == v.String() {
			return v
		}
	}
	return nil
}
