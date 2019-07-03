package fan

type (
	Hu struct {
		Name   string
		Ignore []Fan
		Group  func(hand Hand) []GroupHand
	}
)
