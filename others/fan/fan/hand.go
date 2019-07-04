package fan

type (
	Hand struct {
		Public  []Group
		Private Cards // include the last one
		Last    Card
		ZiMo    bool
	}

	GroupHand struct {
		Groups []Group
		Last   Card
		ZiMo   bool
	}
)

func (h Hand) CardCount() int {
	sum := 0
	for _, _ = range h.Public {
		sum += 3
	}
	for _, count := range h.Private {
		if count < 0 {
			panic("card count can't be negative")
		}
		sum += count
	}
	return sum
}
