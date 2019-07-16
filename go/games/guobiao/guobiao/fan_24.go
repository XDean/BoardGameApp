package guobiao

var (
	ALL_24_FAN = []Fan{
		QI_DUI,
		QI_XING_BU_KAO,
		QING_YI_SE,
		YI_SE_SAN_TONG_SHUN,
		YI_SE_SAN_JIE_GAO,
		QUAN_DA,
		QUAN_ZHONG,
		QUAN_XIAO,
		QUAN_SHUANG_KE,
	}
	QI_DUI = Fan{
		Name:   "七对",
		Fan:    24,
		Ignore: []string{MEN_QIAN_QING.Name, DAN_DIAO_JIANG.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Has(func(g Group) bool {
				return g.Type == GT_QI_DUI
			})
		},
	}

	QI_XING_BU_KAO = Fan{
		Name:   "七星不靠",
		Fan:    24,
		Ignore: []string{MEN_QIAN_QING.Name, DAN_DIAO_JIANG.Name, WU_MEN_QI.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Has(func(g Group) bool {
				return g.Type == GT_QI_XING_BU_KAO
			})
		},
	}

	QING_YI_SE = Fan{
		Name:   "清一色",
		Fan:    24,
		Ignore: []string{WU_ZI.Name},
		Match: func(hand GroupHand) bool {
			for _, t := range TYPE_TBW {
				if hand.Cards.All(TypeIs(t)) {
					return true
				}
			}
			return false
		},
	}

	YI_SE_SAN_TONG_SHUN = Fan{
		Name:   "一色三同顺",
		Fan:    24,
		Ignore: []string{YI_BAN_GAO.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasTriple(func(a Group, b Group, c Group) bool {
				if a.isShunZi() && b.isShunZi() && c.isShunZi() {
					ac := a.Cards.FindMinPointCard()
					bc := b.Cards.FindMinPointCard()
					cc := c.Cards.FindMinPointCard()
					return isEqual(ac, bc, cc)
				}
				return false
			})
		},
	}

	YI_SE_SAN_JIE_GAO = Fan{
		Name:   "一色三节高",
		Fan:    24,
		Ignore: []string{YI_BAN_GAO.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasTriple(func(a Group, b Group, c Group) bool {
				if a.isKeZi() && b.isKeZi() && c.isKeZi() {
					ac := a.Cards.FindMinPointCard()
					bc := b.Cards.FindMinPointCard()
					cc := c.Cards.FindMinPointCard()
					return ac.Type != ZI && isEqual(ac.Type, bc.Type, cc.Type) && isStep(ac.Point, bc.Point, cc.Point)
				}
				return false
			})
		},
	}

	QUAN_DA = Fan{
		Name:   "全大",
		Fan:    24,
		Ignore: []string{WU_ZI.Name, DA_YU_WU.Name},
		Match: func(hand GroupHand) bool {
			return hand.Cards.All(func(c Card) bool {
				return c.Type != ZI && c.Point >= 7
			})
		},
	}

	QUAN_ZHONG = Fan{
		Name:   "全中",
		Fan:    24,
		Ignore: []string{WU_ZI.Name, DUAN_YAO.Name},
		Match: func(hand GroupHand) bool {
			return hand.Cards.All(func(c Card) bool {
				return c.Type != ZI && c.Point >= 4 && c.Point <= 6
			})
		},
	}

	QUAN_XIAO = Fan{
		Name:   "全小",
		Fan:    24,
		Ignore: []string{WU_ZI.Name, XIAO_YU_WU.Name},
		Match: func(hand GroupHand) bool {
			return hand.Cards.All(func(c Card) bool {
				return c.Type != ZI && c.Point <= 3
			})
		},
	}

	QUAN_SHUANG_KE = Fan{
		Name:   "全双刻",
		Fan:    24,
		Ignore: []string{PENG_PENG_HU.Name, DUAN_YAO.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.All(func(g Group) bool {
				return g.isJiang() || g.isKeZi()
			}) && hand.Cards.All(func(card Card) bool {
				return card.Type != ZI && card.Point%2 == 0
			})
		},
	}
)
