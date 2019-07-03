package fan

type (
	Hand struct {
		Public  []Group
		Private Cards
		Last    Card
		ZiMo    bool
	}

	GroupHand struct {
		Groups []Group
		Last   Card
		ZiMo   bool
	}
)
