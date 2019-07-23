package guobiao

var (
	ALL_12_FAN = []Fan{
		QUAN_BU_KAO,
		ZU_HE_LONG,
		XIAO_YU_WU,
		DA_YU_WU,
		SAN_FENG_KE,
	}
	QUAN_BU_KAO = Fan{
		Name:   "全不靠",
		Fan:    12,
		Ignore: []string{WU_MEN_QI.Name, MEN_QIAN_QING.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Has(func(g Group) bool {
				return g.Type == GT_QUAN_BU_KAO
			})
		},
	}

	ZU_HE_LONG = Fan{
		Name: "组合龙",
		Fan:  12,
		Match: func(hand GroupHand) bool {
			return hand.Groups.Has(func(g Group) bool {
				return g.Type == GT_ZU_HE_LONG
			})
		},
	}

	DA_YU_WU = Fan{
		Name:   "大于五",
		Fan:    12,
		Ignore: []string{WU_ZI.Name},
		Match: func(hand GroupHand) bool {
			return hand.Cards.All(func(c Card) bool {
				return c.Type != ZI && c.Point > 5
			})
		},
	}

	XIAO_YU_WU = Fan{
		Name:   "小于五",
		Fan:    12,
		Ignore: []string{WU_ZI.Name},
		Match: func(hand GroupHand) bool {
			return hand.Cards.All(func(c Card) bool {
				return c.Type != ZI && c.Point < 5
			})
		},
	}

	SAN_FENG_KE = Fan{
		Name:   "三风刻",
		Fan:    12,
		Ignore: []string{YAO_JIU_KE.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Find(func(g Group) bool {
				return g.isKeZi() && g.Cards.FindMinPointCard().isFeng()
			}).Size() == 3
		},
	}
)
