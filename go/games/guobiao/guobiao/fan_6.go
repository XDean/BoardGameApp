package guobiao

import "sort"

var (
	ALL_6_FAN = []Fan{
		PENG_PENG_HU,
		HUN_YI_SE,
		QUAN_QIU_REN,
		SAN_SE_SAN_BU_GAO,
		WU_MEN_QI,
		SHUANG_AN_GANG,
		SHUANG_JIAN_KE,
	}
	PENG_PENG_HU = Fan{
		Name: "碰碰胡",
		Fan:  6,
		Match: func(hand GroupHand) bool {
			return !hand.Groups.HasGroup(func(g Group) bool {
				return !g.isKeZi() && !g.isJiang()
			})
		},
	}

	HUN_YI_SE = Fan{
		Name: "混一色",
		Fan:  6,
		Match: func(hand GroupHand) bool {
			for _, t := range TYPE_TBW {
				if hand.Cards.Find(TypeIs(t).Or(TypeIs(ZI))).Size() == hand.Cards.Size() {
					return true
				}
			}
			return false
		},
	}

	QUAN_QIU_REN = Fan{
		Name:   "全求人",
		Fan:    6,
		Ignore: []string{DAN_DIAO_JIANG.Name},
		Match: func(hand GroupHand) bool {
			if hand.ZiMo {
				return false
			}
			if !hand.FindLastGroup().HasGroup(func(g Group) bool {
				return g.isJiang()
			}) {
				return false
			}
			return !hand.Groups.HasGroup(func(g Group) bool {
				return !g.isChiPengMing() && !g.isJiang()
			})
		},
	}

	SAN_SE_SAN_BU_GAO = Fan{
		Name: "三色三步高",
		Fan:  6,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroupTriple(func(a Group, b Group, c Group) bool {
				if a.isShunZi() && b.isShunZi() && c.isShunZi() {
					am := a.Cards.FindMinPointCard()
					bm := b.Cards.FindMinPointCard()
					cm := c.Cards.FindMinPointCard()
					if am.Type != bm.Type && bm.Type != cm.Type && am.Type != cm.Type {
						return isStep(am.Point, bm.Point, cm.Point)
					}
				}
				return false
			})
		},
	}
	WU_MEN_QI = Fan{
		Name: "五门齐",
		Fan:  6,
		Match: func(hand GroupHand) bool {
			return hand.Cards.Has(TypeIs(TIAO)) &&
				hand.Cards.Has(TypeIs(BING)) &&
				hand.Cards.Has(TypeIs(WAN)) &&
				hand.Cards.Has(func(c Card) bool {
					return c.isZFB()
				}) &&
				hand.Cards.Has(func(c Card) bool {
					return c.isFeng()
				})
		},
	}

	SHUANG_AN_GANG = Fan{
		Name:   "双暗杠",
		Fan:    6,
		Ignore: []string{AN_GANG.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.FindGroup(func(g Group) bool {
				return g.Type == GT_AN_GANG
			}).Size() == 2
		},
	}

	SHUANG_JIAN_KE = Fan{
		Name:   "双箭刻",
		Fan:    6,
		Ignore: []string{JIAN_KE.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.FindGroup(func(g Group) bool {
				return g.isKeZi() && g.Cards.FindMinPointCard().isZFB()
			}).Size() == 2
		},
	}
)
