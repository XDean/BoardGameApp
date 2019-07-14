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

func (gs Groups) HasPair(match func(Group, Group) bool) bool {
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

func (gs Groups) HasTriple(match func(Group, Group, Group) bool) bool {
	for i1, g1 := range gs {
		for i2 := i1 + 1; i2 < len(gs); i2++ {
			g2 := gs[i2]
			for i3 := i2 + 1; i3 < len(gs); i3++ {
				g3 := gs[i3]
				if match(g1, g2, g3) {
					return true
				}
			}
		}
	}
	return false
}

func (gs Groups) HasQuad(match func(Group, Group, Group, Group) bool) bool {
	if gs.Size() < 4 {
		return false
	}
	return match(gs[0], gs[1], gs[2], gs[3])
}

func (gs Groups) Has(match func(Group) bool) bool {
	for _, g := range gs {
		if match(g) {
			return true
		}
	}
	return false
}

func (gs Groups) All(match func(Group) bool) bool {
	return gs.Find(match).Size() == gs.Size()
}

func (gs Groups) Find(match func(Group) bool) Groups {
	result := make(Groups, 0)
	for _, g := range gs {
		if match(g) {
			result = append(result, g)
		}
	}
	return result
}

func (gs Groups) FindPair(match func(Group, Group) bool) (Groups, Groups) {
	left := make(Groups, 0)
	right := make(Groups, 0)
	for i1, g1 := range gs {
		for i2 := i1 + 1; i2 < len(gs); i2++ {
			g2 := gs[i2]
			if match(g1, g2) {
				left = append(left, g1)
				right = append(right, g2)
			}
		}
	}
	return left, right
}

func (gs Groups) Size() int {
	return len(gs)
}

func (h GroupHand) FindLastGroup() Groups {
	result := make(Groups, 0)
	for _, g := range h.Groups {
		if g.Cards.Find(CardIs(h.Last)).Size() > 0 {
			if !g.Type.Public {
				result = append(result, g)
			}
		}
	}
	return result
}
