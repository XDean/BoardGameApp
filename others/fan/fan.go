package fan

type Fan struct {
	Name   string
	Fan    int
	Match  func(hand Hand) bool
	Ignore []string
}

type ExtraFan struct {
	Name   string
	Fan    int
	Ignore []string
}
