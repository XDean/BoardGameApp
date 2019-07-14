package guobiao

var (
	ALL_4_FAN = []Fan{
		BU_QIU_REN,
		QUAN_DAI_YAO,
		SHUANG_MING_GANG,
	}

	HU_JUE_ZHANG = ExtraFan{
		Name: "和绝张",
		Fan:  4,
	}

	BU_QIU_REN = Fan{
		Name: "不求人",
		Fan:  4,
		Match: func(hand GroupHand) bool {
			return hand.ZiMo && !hand.Groups.HasGroup(func(g Group) bool {
				return g.isChiPengMing()
			})
		},
	}

	QUAN_DAI_YAO = Fan{
		Name: "全带幺",
		Fan:  4,
		Match: func(hand GroupHand) bool {
			return !hand.Groups.HasGroup(func(g Group) bool {
				return !g.Cards.Has(func(c Card) bool {
					return c.Point == 1 || c.Point == 9 || c.Type == ZI
				})
			})
		},
	}

	SHUANG_MING_GANG = Fan{
		Name: "双明杠",
		Fan:  4,
		Match: func(hand GroupHand) bool {
			return hand.Groups.FindGroup(func(g Group) bool {
				return g.Type == GT_MING_GANG
			}).Size() == 2
		},
	}
)
