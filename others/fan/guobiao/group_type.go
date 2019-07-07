package guobiao

type (
	GroupType struct {
		Name      string
		Public    bool
		CardCount int
	}
)

func (gt GroupType) Find(cards Cards, card Card) (bool, Group, Cards) {
	switch gt {
	case GT_CHI:
		return chiFind(cards, card)
	case GT_PENG:
		return pengFind(cards, card)
	case GT_MING_GANG:
		return mgFind(cards, card)
	case GT_AN_GANG:
		return agFind(cards, card)
	case GT_KE:
		return keFind(cards, card)
	case GT_SHUN:
		return shunFind(cards, card)
	case GT_ZU_HE_LONG:
		return zhlFind(cards, card)
	case GT_QUAN_BU_KAO:
		return qbkFind(cards, card)
	case GT_QI_XING_BU_KAO:
		return qxbkFind(cards, card)
	case GT_QI_DUI:
		return qiduiFind(cards, card)
	case GT_SHI_SAN_YAO:
		return ssyFind(cards, card)
	}
	panic("never happen")
}

var (
	GT_CHI = GroupType{
		Public:    true,
		CardCount: 3,
	}
	GT_PENG = GroupType{
		Public:    true,
		CardCount: 3,
	}
	GT_MING_GANG = GroupType{
		Public:    true,
		CardCount: 4,
	}
	GT_AN_GANG = GroupType{
		Public:    true,
		CardCount: 4,
	}

	GT_KE = GroupType{
		Public:    false,
		CardCount: 3,
	}
	GT_SHUN = GroupType{
		Public:    false,
		CardCount: 3,
	}

	GT_ZU_HE_LONG = GroupType{
		Public:    false,
		CardCount: 9,
	}

	GT_QI_DUI = GroupType{
		Public:    false,
		CardCount: 14,
	}

	GT_QUAN_BU_KAO = GroupType{
		Public:    false,
		CardCount: 14,
	}

	GT_QI_XING_BU_KAO = GroupType{
		Public:    false,
		CardCount: 14,
	}

	GT_SHI_SAN_YAO = GroupType{
		Public:    false,
		CardCount: 14,
	}
)

func chiFind(cards Cards, card Card) (bool, Group, Cards) {
	b, group, i := shunFind(cards, card)
	group.Type = GT_CHI
	return b, group, i
}

func pengFind(cards Cards, card Card) (bool, Group, Cards) {
	b, group, i := keFind(cards, card)
	group.Type = GT_PENG
	return b, group, i
}

func mgFind(cards Cards, card Card) (bool, Group, Cards) {
	left := cards.Copy()
	use := Cards{}
	left.MoveTo(use, card, 4)
	if left.IsValid() {
		return true, Group{
			Type:  GT_MING_GANG,
			Cards: use,
		}, left
	}
	return false, Group{}, nil
}

func agFind(cards Cards, card Card) (bool, Group, Cards) {
	b, group, i := mgFind(cards, card)
	group.Type = GT_AN_GANG
	return b, group, i
}

func keFind(cards Cards, card Card) (bool, Group, Cards) {
	left := cards.Copy()
	use := Cards{}
	left.MoveTo(use, card, 3)
	if left.IsValid() {
		return true, Group{
			Type:  GT_KE,
			Cards: use,
		}, left
	}
	return false, Group{}, nil
}

func shunFind(cards Cards, card Card) (bool, Group, Cards) {
	if !card.isTBW() || card.Point > 7 {
		return false, Group{}, nil
	}
	left := cards.Copy()
	use := Cards{}
	left.MoveTo(use, card, 1)
	left.MoveTo(use, card.Next(1), 1)
	left.MoveTo(use, card.Next(2), 1)
	if left.IsValid() {
		return true, Group{
			Type:  GT_SHUN,
			Cards: use,
		}, left
	}
	return false, Group{}, nil
}

func zhlFind(cards Cards, card Card) (bool, Group, Cards) {
	if !card.isTBW() || cards.Size() < 9 {
		return false, Group{}, nil
	}
	for tiao := 1; tiao <= 3; tiao++ {
		for bing := 1; bing <= 3; bing++ {
			for wan := 1; wan <= 3; wan++ {
				if tiao != bing && tiao != wan && bing != wan {
					t := Card{Type: TIAO, Point: tiao}
					b := Card{Type: BING, Point: bing}
					w := Card{Type: WAN, Point: wan}
					left := cards.Copy()
					use := cards
					left.MoveTo(use, t, 1)
					left.MoveTo(use, t.Next(3), 1)
					left.MoveTo(use, t.Next(6), 1)
					left.MoveTo(use, b, 1)
					left.MoveTo(use, b.Next(3), 1)
					left.MoveTo(use, b.Next(6), 1)
					left.MoveTo(use, w, 1)
					left.MoveTo(use, w.Next(3), 1)
					left.MoveTo(use, w.Next(6), 1)
					if left[card] > 0 && left.IsValid() {
						return true, Group{
							Type:  GT_ZU_HE_LONG,
							Cards: use,
						}, left
					}
				}
			}
		}
	}
	return false, Group{}, nil
}

func qiduiFind(cards Cards, card Card) (bool, Group, Cards) {
	if cards.Size() != 14 {
		return false, Group{}, nil
	}
	for _, count := range cards {
		if count != 2 && count != 4 {
			return false, Group{}, nil
		}
	}
	return true, Group{
		Type:  GT_QI_DUI,
		Cards: cards.Copy(),
	}, Cards{}
}

func qbkFind(cards Cards, card Card) (bool, Group, Cards) {
	if cards.Size() != 14 {
		return false, Group{}, nil
	}
	left := cards.Copy()
	use := Cards{}
	for card, count := range cards {
		if count > 1 {
			return false, Group{}, nil
		}
		if !card.isZi() {
			if use.Find(PointNear(card.Point, 2)).Find(TypeIs(card.Type)).Size() != 0 {
				return false, Group{}, nil
			}
			if use.Find(PointIs(card.Point)).Size() != 0 {
				return false, Group{}, nil
			}
		}
		left.MoveTo(use, card, 1)
	}
	return true, Group{
		Type:  GT_QUAN_BU_KAO,
		Cards: use,
	}, left
}

func qxbkFind(cards Cards, card Card) (bool, Group, Cards) {
	if ok, group, left := qbkFind(cards, card); ok {
		if group.Cards.Find(TypeIs(ZI)).Size() == 7 {
			group.Type = GT_QI_XING_BU_KAO
			return true, group, left
		}
	}
	return false, Group{}, nil
}

func ssyFind(cards Cards, card Card) (bool, Group, Cards) {
	if cards.Size() != 14 {
		return false, Group{}, nil
	}
	left := cards.Copy()
	use := Cards{}
	var extra bool
	for card, count := range cards {
		if count > 1 {
			if extra || count > 2 {
				return false, Group{}, nil
			} else {
				extra = true
			}
		}
		if !card.isZi() && card.Point != 1 && card.Point != 9 {
			return false, Group{}, nil
		}
		left.MoveTo(use, card, count)
	}
	return true, Group{
		Type:  GT_SHI_SAN_YAO,
		Cards: use,
	}, left
}
