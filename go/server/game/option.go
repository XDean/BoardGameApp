package game

type (
	Option struct {
		Id     string
		Name   string
		Type   OptionType
		Domain interface{}
	}

	OptionType int

	Range struct {
		Min int
		Max int
	}
)
