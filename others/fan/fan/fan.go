package fan

type (
	Fan struct {
		Name   string
		Fan    int
		Ignore []string
		Match  func(hand Hand) bool
	}
)
