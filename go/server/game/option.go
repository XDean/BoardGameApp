package game

type (
	Option struct {
		Id   string
		Name string
		Type OptionType
		Data interface{}
	}

	OptionType int

	Range struct {
		Min int
		Max int
	}
)

func (r Range) Contain(i int) bool {
	return i >= r.Min && i <= r.Max
}
