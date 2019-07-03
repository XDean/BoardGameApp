package fan

type (
	GroupType struct {
		Public    bool
		CardCount int
		Find      func(cards Cards, card Card) (bool, Group, Cards)
	}
)

var (
	CHI = GroupType{
		Public:    true,
		CardCount: 3,
		Find:      noFind,
	}
	PENG = GroupType{
		Public:    true,
		CardCount: 3,
		Find:      noFind,
	}
	MING_GANG = GroupType{
		Public:    true,
		CardCount: 4,
		Find:      noFind,
	}
	AN_GANG = GroupType{
		Public:    true,
		CardCount: 4,
		Find:      noFind,
	}

	KE = GroupType{
		Public:    false,
		CardCount: 3,
		Find:      keFind,
	}
	SHUN = GroupType{
		Public:    false,
		CardCount: 3,
		Find:      shunFind,
	}

	ZU_HE_LONG = GroupType{
		Public:    false,
		CardCount: 9,
	}

	QI_DUI = GroupType{
		Public:    false,
		CardCount: 14,
	}

	QUAN_BU_KAO = GroupType{
		Public:    false,
		CardCount: 14,
	}

	QI_XING_BU_KAO = GroupType{
		Public:    false,
		CardCount: 14,
	}

	SHI_SAN_YAO = GroupType{
		Public:    false,
		CardCount: 14,
	}
)

func noFind(cards Cards, card Card) (bool, Group, Cards) {
	return false, Group{}, nil
}

func keFind(cards Cards, card Card) (bool, Group, Cards) {
	left := cards.Copy()
	use := Cards{}
	left.MoveTo(use, card, 3)
	if left.IsValid() {
		return true, Group{
			Type:  KE,
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
			Type:  SHUN,
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
							Type:  ZU_HE_LONG,
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
		Type:  QI_DUI,
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
		Type:  QUAN_BU_KAO,
		Cards: use,
	}, left
}

func qxbkFind(cards Cards, card Card) (bool, Group, Cards) {
	if ok, group, left := qbkFind(cards, card); ok {
		if group.Cards.Find(TypeIs(ZI)).Size() == 7 {
			group.Type = QI_XING_BU_KAO
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
			if extra {
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
		Type:  SHI_SAN_YAO,
		Cards: use,
	}, left
}
