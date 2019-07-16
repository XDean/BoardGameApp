package guobiao

var (
	ALL_88_FAN = []Fan{
		DA_SI_XI,
		DA_SAN_YUAN,
		SHI_SAN_YAO,
		LV_YI_SE,
		JIU_LIAN_BAO_DENG,
		LIAN_QI_DUI,
		SI_GANG,
	}
	DA_SI_XI = Fan{
		Name:   "大四喜",
		Fan:    88,
		Ignore: []string{SAN_FENG_KE.Name, YAO_JIU_KE.Name, PENG_PENG_HU.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Find(func(g Group) bool {
				return g.isKeZi() && g.Cards.Any().isFeng()
			}).Size() == 4
		},
	}
	DA_SAN_YUAN = Fan{
		Name:   "大三元",
		Fan:    88,
		Ignore: []string{SHUANG_JIAN_KE.Name, JIAN_KE.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Find(func(g Group) bool {
				return g.isKeZi() && g.Cards.Any().isZFB()
			}).Size() == 3
		},
	}
	SHI_SAN_YAO = Fan{
		Name:   "十三幺",
		Fan:    88,
		Ignore: []string{WU_MEN_QI.Name, MEN_QIAN_QING.Name, DAN_DIAO_JIANG.Name, HUN_YAO_JIU.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Has(func(g Group) bool {
				return g.Type == GT_SHI_SAN_YAO
			})
		},
	}
	cardLvYiSe = []Card{{Type: TIAO, Point: 2}, {Type: TIAO, Point: 3}, {Type: TIAO, Point: 4},
		{Type: TIAO, Point: 6}, {Type: TIAO, Point: 8}, {Type: ZI, Point: Z_FA}}
	LV_YI_SE = Fan{
		Name: "绿一色",
		Fan:  88,
		Match: func(hand GroupHand) bool {
			return hand.Cards.All(CardIn(cardLvYiSe))
		},
	}
	JIU_LIAN_BAO_DENG = Fan{
		Name:   "九莲宝灯",
		Fan:    88,
		Ignore: []string{QING_YI_SE.Name, YAO_JIU_KE.Name, MEN_QIAN_QING.Name},
		Match: func(hand GroupHand) bool {
			if hand.Groups.Has(func(g Group) bool {
				return g.Type.Public
			}) {
				return false
			}
			t := hand.Cards.Any().Type
			if t == ZI {
				return false
			}
			copy := hand.Cards.Copy()
			use := Cards{}
			for i := 1; i <= 9; i++ {
				if i == 1 || i == 9 {
					copy.MoveTo(use, Card{Type: t, Point: i}, 3)
				} else {
					copy.MoveTo(use, Card{Type: t, Point: i}, 1)
				}
			}
			if copy.IsValid() {
				return copy.Any().Type == t
			}
			return false
		},
	}
	LIAN_QI_DUI = Fan{
		Name:   "连七对",
		Fan:    88,
		Ignore: []string{QING_YI_SE.Name, QI_DUI.Name, DAN_DIAO_JIANG.Name, MEN_QIAN_QING.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Has(func(g Group) bool {
				if g.Type != GT_QI_DUI {
					return false
				}
				min := g.Cards.FindMinPointCard()
				for i := 0; i < 7; i++ {
					if !g.Cards.HasCard(Card{Type: min.Type, Point: min.Point + i}) {
						return false
					}
				}
				return min.Type != ZI
			})
		},
	}
	SI_GANG = Fan{
		Name:   "四杠",
		Fan:    88,
		Ignore: []string{SHUANG_AN_GANG.Name, DAN_DIAO_JIANG.Name, PENG_PENG_HU.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Find(func(g Group) bool {
				return g.isGang()
			}).Size() == 4
		},
	}
)
