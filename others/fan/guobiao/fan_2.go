package guobiao

var (
	ALL_2_FAN = []Fan{
		MEN_QIAN_QING,
		DUAN_YAO,
		PING_HU,
		JIAN_KE,
		SI_GUI_YI,
		SHUANG_TONG_KE,
		SHUANG_AN_KE,
		AN_GANG,
	}
	MEN_QIAN_QING = Fan{
		Name: "门前清",
		Fan:  2,
		Match: func(hand GroupHand) bool {
			return !hand.ZiMo && !hand.Groups.HasGroup(func(g Group) bool {
				return g.isChiPengMing()
			})
		},
	}
	DUAN_YAO = Fan{
		Name: "断幺",
		Fan:  2,
		Match: func(hand GroupHand) bool {
			return hand.Cards.Find(PointIs(1).
				Or(PointIs(9).Or(TypeIs(ZI)))).Size() == 0
		},
	}
	PING_HU = Fan{
		Name:   "平胡",
		Fan:    2,
		Ignore: []string{WU_ZI.Name},
		Match: func(hand GroupHand) bool {
			return !hand.Groups.HasGroup(func(g Group) bool {
				return !(g.Type == GT_JIANG || g.Type == GT_SHUN)
			})
		},
	}
	JIAN_KE = Fan{
		Name: "箭刻",
		Fan:  2,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroup(func(g Group) bool {
				return g.isKeZi() && g.Cards.FindMinPointCard().isZFB()
			})
		},
	}
	SI_GUI_YI = Fan{
		Name: "四归一",
		Fan:  2,
		Match: func(hand GroupHand) bool {
			return hand.Cards.Has(func(c Card) bool {
				if hand.Cards[c] == 4 {
					return !hand.Groups.HasGroup(func(g Group) bool {
						return g.isGang() && g.Cards.HasCard(c)
					})
				} else {
					return false
				}
			})
		},
	}
	SHUANG_TONG_KE = Fan{
		Name: "双同刻",
		Fan:  2,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroupPair(func(a Group, b Group) bool {
				return a.isKeZi() && b.isKeZi() && a.Cards.FindMinPointCard().Point == b.Cards.FindMinPointCard().Point
			})
		},
	}
	SHUANG_AN_KE = Fan{
		Name: "双暗刻",
		Fan:  2,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroupPair(func(a Group, b Group) bool {
				return a.isAnKe() && b.isAnKe()
			})
		},
	}
	AN_GANG = Fan{
		Name: "暗杠",
		Fan:  2,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroup(func(g Group) bool {
				return g.Type == GT_AN_GANG
			})
		},
	}
)
