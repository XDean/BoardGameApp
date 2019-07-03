package fan

type (
	Fan struct {
		Name   string
		Fan    int
		Ignore []interface{}
		Match  func(hand Hand) bool
	}

	ExtraFan struct {
		Name   string
		Fan    int
		Ignore []interface{}
	}
)
