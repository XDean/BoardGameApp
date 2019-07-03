package fan

type (
	Hand struct {
		Public  Cards
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
