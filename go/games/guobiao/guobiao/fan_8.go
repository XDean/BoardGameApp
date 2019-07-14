package guobiao

var (
	ALL_8_FAN = []Fan{
		HUA_LONG,
		SAN_SE_SAN_TONG_SHUN,
		SAN_SE_SAN_JIE_GAO,
		TUI_BU_DAO,
	}
	HUA_LONG = Fan{
		Name: "花龙",
		Fan:  8,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasTriple(func(a Group, b Group, c Group) bool {
				if a.isShunZi() && b.isShunZi() && c.isShunZi() {
					am := a.Cards.FindMinPointCard()
					bm := b.Cards.FindMinPointCard()
					cm := c.Cards.FindMinPointCard()
					if isTBW(am.Type, bm.Type, cm.Type) {
						return isStep3(am.Point, bm.Point, cm.Point)
					}
				}
				return false
			})
		},
	}
	SAN_SE_SAN_TONG_SHUN = Fan{
		Name: "三色三同顺",
		Fan:  8,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasTriple(func(a Group, b Group, c Group) bool {
				if a.isShunZi() && b.isShunZi() && c.isShunZi() {
					am := a.Cards.FindMinPointCard()
					bm := b.Cards.FindMinPointCard()
					cm := c.Cards.FindMinPointCard()
					if am.Type != bm.Type && bm.Type != cm.Type && am.Type != cm.Type {
						return isEqual(am.Point, bm.Point, cm.Point)
					}
				}
				return false
			})
		},
	}
	SAN_SE_SAN_JIE_GAO = Fan{
		Name: "三色三节高",
		Fan:  8,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasTriple(func(a Group, b Group, c Group) bool {
				if a.isKeZi() && b.isKeZi() && c.isKeZi() {
					am := a.Cards.FindMinPointCard()
					bm := b.Cards.FindMinPointCard()
					cm := c.Cards.FindMinPointCard()
					if isTBW(am.Type, bm.Type, cm.Type) {
						return isStep(am.Point, bm.Point, cm.Point)
					}
				}
				return false
			})
		},
	}

	WU_FAN_HU = Fan{
		Name: "无番和",
		Fan:  8,
		Match: func(hand GroupHand) bool {
			return false
		},
	}

	cardTuiBuDao = []Card{{Type: BING, Point: 1}, {Type: BING, Point: 2}, {Type: BING, Point: 3}, {Type: BING, Point: 4}, {Type: BING, Point: 5},
		{Type: BING, Point: 8}, {Type: BING, Point: 9}, {Type: TIAO, Point: 2}, {Type: TIAO, Point: 4}, {Type: TIAO, Point: 5}, {Type: TIAO, Point: 6}, {Type: TIAO, Point: 8},
		{Type: TIAO, Point: 9}, {Type: ZI, Point: Z_BAI}}
	TUI_BU_DAO = Fan{
		Name:   "推不到",
		Fan:    8,
		Ignore: []string{QUE_YI_MEN.Name},
		Match: func(hand GroupHand) bool {
			return hand.Cards.All(CardIn(cardTuiBuDao))
		},
	}

	MIAO_SHOU_HUI_CHUN = ExtraFan{
		Name:   "妙手回春",
		Fan:    8,
		Ignore: []string{ZI_MO.Name},
	}

	HAI_DI_LAO_YUE = ExtraFan{
		Name: "海底捞月",
		Fan:  8,
	}

	GANG_SHANG_KAI_HUA = ExtraFan{
		Name:   "杠上开花",
		Fan:    8,
		Ignore: []string{ZI_MO.Name},
	}

	QIANG_GANG_HU = ExtraFan{
		Name:   "抢杠和",
		Fan:    8,
		Ignore: []string{HU_JUE_ZHANG.Name},
	}
)
