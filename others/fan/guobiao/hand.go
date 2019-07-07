package guobiao

type (
	Groups []Group

	Hand struct {
		Public  Groups
		Private Cards // include the last one
		Last    Card
		ZiMo    bool
	}

	GroupHand struct {
		Groups Groups
		Cards  Cards
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

func (gs Groups) HasGroupPair(match func(Group, Group) bool) bool {
	for i1, g1 := range gs {
		for i2 := i1 + 1; i2 < len(gs); i2++ {
			g2 := gs[i2]
			if match(g1, g2) {
				return true
			}
		}
	}
	return false
}

func (gs Groups) HasGroup(match func(Group) bool) bool {
	for _, g := range gs {
		if match(g) {
			return true
		}
	}
	return false
}

func (h GroupHand) FindLastGroup() Groups {
	result := make(Groups, 0)
	for _, g := range h.Groups {
		if g.Cards.Find(CardIs(h.Last)).Size() > 0 {
			result = append(result, g)
		}
	}
	return result
}
