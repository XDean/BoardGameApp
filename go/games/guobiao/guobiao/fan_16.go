package guobiao

var (
	ALL_16_FAN = []Fan{
		QING_LONG,
		YI_SE_SAN_BU_GAO,
		SAN_SE_SHUANG_LONG_HUI,
		QUAN_DAI_WU,
		SAN_TONG_KE,
		SAN_AN_KE,
	}
	QING_LONG = Fan{
		Name: "清龙",
		Fan:  16,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasTriple(func(a Group, b Group, c Group) bool {
				if a.isShunZi() && b.isShunZi() && c.isShunZi() {
					am := a.Cards.FindMinPointCard()
					bm := b.Cards.FindMinPointCard()
					cm := c.Cards.FindMinPointCard()
					if isEqual(am.Type, bm.Type, cm.Type) {
						return isStep3(am.Point, bm.Point, cm.Point)
					}
				}
				return false
			})
		},
	}
	YI_SE_SAN_BU_GAO = Fan{
		Name: "一色三步高",
		Fan:  16,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasTriple(func(a Group, b Group, c Group) bool {
				if a.isShunZi() && b.isShunZi() && c.isShunZi() {
					am := a.Cards.FindMinPointCard()
					bm := b.Cards.FindMinPointCard()
					cm := c.Cards.FindMinPointCard()
					if isEqual(am.Type, bm.Type, cm.Type) {
						return isStep(am.Point, bm.Point, cm.Point)
					}
				}
				return false
			})
		},
	}
	SAN_SE_SHUANG_LONG_HUI = Fan{
		Name:   "三色双龙会",
		Fan:    16,
		Ignore: []string{XI_XIANG_FENG.Name, LAO_SHAO_FU.Name, WU_ZI.Name, PING_HU.Name},
		Match: func(hand GroupHand) bool {
			jiang := hand.Groups.Find(func(g Group) bool {
				return g.isJiang()
			})
			left, _ := hand.Groups.FindPair(func(a Group, b Group) bool {
				if a.isShunZi() && b.isShunZi() {
					aMax := a.Cards.FindMaxPointCard()
					bMax := b.Cards.FindMaxPointCard()
					if aMax.Type == bMax.Type && Abs(aMax.Point-bMax.Point) == 6 {
						return true
					}
				}
				return false
			})
			if jiang.Size() != 1 || left.Size() != 2 {
				return false
			}
			return isTBW(jiang[0].Cards.Any().Type, left[0].Cards.Any().Type, left[1].Cards.Any().Type)
		},
	}
	QUAN_DAI_WU = Fan{
		Name:   "全带五",
		Fan:    16,
		Ignore: []string{DUAN_YAO.Name},
		Match: func(hand GroupHand) bool {
			return !hand.Groups.Has(func(g Group) bool {
				return !g.Cards.Has(PointIs(5))
			})
		},
	}
	SAN_TONG_KE = Fan{
		Name: "三同刻",
		Fan:  16,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasTriple(func(a Group, b Group, c Group) bool {
				if a.isKeZi() && b.isKeZi() && c.isKeZi() {
					ac := a.Cards.Any()
					bc := b.Cards.Any()
					cc := c.Cards.Any()
					return isTBW(ac.Type, bc.Type, cc.Type) && isEqual(ac.Point, bc.Point, cc.Point)
				}
				return false
			})
		},
	}
	SAN_AN_KE = Fan{
		Name:   "三暗刻",
		Fan:    16,
		Ignore: []string{SHUANG_AN_KE.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Find(func(a Group) bool {
				return a.isAnKe()
			}).Size() == 3
		},
	}
)
