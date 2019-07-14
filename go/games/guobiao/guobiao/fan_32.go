package guobiao

var (
	ALL_32_FAN = []Fan{
		YI_SE_SI_BU_GAO,
		HUN_YAO_JIU,
		SAN_GANG,
	}
	YI_SE_SI_BU_GAO = Fan{
		Name:   "一色四步高",
		Fan:    32,
		Ignore: []string{LIAN_LIU.Name, LAO_SHAO_FU.Name, YI_SE_SAN_BU_GAO.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasQuad(func(a, b, c, d Group) bool {
				if a.isShunZi() && b.isShunZi() && c.isShunZi() && d.isShunZi() {
					am := a.Cards.FindMinPointCard()
					bm := b.Cards.FindMinPointCard()
					cm := c.Cards.FindMinPointCard()
					dm := d.Cards.FindMinPointCard()
					if isEqual(am.Type, bm.Type, cm.Type, dm.Type) {
						return isStep(am.Point, bm.Point, cm.Point, dm.Point)
					}
				}
				return false
			})
		},
	}

	HUN_YAO_JIU = Fan{
		Name:   "混幺九",
		Fan:    32,
		Ignore: []string{PENG_PENG_HU.Name, QUAN_DAI_YAO.Name, YAO_JIU_KE.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.All(func(g Group) bool {
				return g.isJiang() || g.isKeZi()
			}) && hand.Cards.All(func(card Card) bool {
				return card.Type == ZI || card.Point == 1 || card.Point == 9
			})
		},
	}

	SAN_GANG = Fan{
		Name: "三杠",
		Fan:  32,
		Match: func(hand GroupHand) bool {
			return hand.Groups.Find(func(g Group) bool {
				return g.isGang()
			}).Size() == 3
		},
	}
)
