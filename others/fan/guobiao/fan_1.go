package guobiao

var (
	ALL_1_FAN = []Fan{
		QUE_YI_MEN,
		WU_ZI,
		YI_BAN_GAO,
		LIAN_LIU,
		LAO_SHAO_FU,
		XI_XIANG_FENG,
		YAO_JIU_KE,
		MING_GANG,
		BIAN_ZHANG,
		KAN_ZHANG,
		DAN_DIAO_JIANG,
		ZI_MO,
	}
	QUE_YI_MEN = Fan{
		Name: "缺一门",
		Fan:  1,
		Match: func(hand GroupHand) bool {
			for _, t := range TYPE_TBW {
				if hand.Cards.Find(TypeIs(t)).Size() == 0 {
					return true
				}
			}
			return false
		},
	}
	WU_ZI = Fan{
		Name: "无字",
		Fan:  1,
		Match: func(hand GroupHand) bool {
			if hand.Cards.Find(TypeIs(ZI)).Size() == 0 {
				return true
			}
			return false
		},
	}
	YI_BAN_GAO = Fan{
		Name: "一般高",
		Fan:  1,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroupPair(func(a Group, b Group) bool {
				if a.isShunZi() && b.isShunZi() {
					if a.Cards.Equals(b.Cards) {
						return true
					}
				}
				return false
			})
		},
	}
	LIAN_LIU = Fan{
		Name: "连六",
		Fan:  1,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroupPair(func(a Group, b Group) bool {
				if a.isShunZi() && b.isShunZi() {
					aMax := a.Cards.FindMaxPoint()
					bMax := b.Cards.FindMaxPoint()
					if aMax.Type == bMax.Type && Abs(aMax.Point-bMax.Point) == 3 {
						return true
					}
				}
				return false
			})
		},
	}
	LAO_SHAO_FU = Fan{
		Name: "老少副",
		Fan:  1,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroupPair(func(a Group, b Group) bool {
				if a.isShunZi() && b.isShunZi() {
					aMax := a.Cards.FindMaxPoint()
					bMax := b.Cards.FindMaxPoint()
					if aMax.Type == bMax.Type && Abs(aMax.Point-bMax.Point) == 6 {
						return true
					}
				}
				return false
			})
		},
	}
	XI_XIANG_FENG = Fan{
		Name: "喜相逢",
		Fan:  1,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroupPair(func(a Group, b Group) bool {
				if a.isShunZi() && b.isShunZi() {
					aMax := a.Cards.FindMaxPoint()
					bMax := b.Cards.FindMaxPoint()
					if aMax.Type != bMax.Type && Abs(aMax.Point-bMax.Point) == 0 {
						return true
					}
				}
				return false
			})
		},
	}
	YAO_JIU_KE = Fan{
		Name: "幺九刻",
		Fan:  1,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroup(func(a Group) bool {
				if a.isKeZi() {
					aMax := a.Cards.FindMaxPoint()
					if aMax.Point == 1 || aMax.Point == 9 {
						return true
					}
				}
				return false
			})
		},
	}
	MING_GANG = Fan{
		Name: "明杠",
		Fan:  1,
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasGroup(func(a Group) bool {
				return a.isGang()
			})
		},
	}
	BIAN_ZHANG = Fan{
		Name: "边张",
		Fan:  1,
		Match: func(hand GroupHand) bool {
			if hand.Last.Point != 3 && hand.Last.Point != 7 {
				return false
			}
			return !hand.FindLastGroup().HasGroupPair(func(a Group, b Group) bool {
				if a.Cards.Find(CardIs(hand.Last)).Size() > 0 && b.Cards.Find(CardIs(hand.Last)).Size() > 0 {
					aMax := a.Cards.FindMaxPoint()
					bMax := b.Cards.FindMaxPoint()
					if Abs(aMax.Point-bMax.Point) == 2 {
						return true
					}
				}
				return false
			})
		},
	}
	KAN_ZHANG = Fan{
		Name: "坎张",
		Fan:  1,
		Match: func(hand GroupHand) bool {

			return hand.FindLastGroup().HasGroup(func(g Group) bool {
				return g.isShunZi() && g.Cards.FindMinPointCard().Point == hand.Last.Point-1
			}) && !hand.Groups.HasGroupPair(func(a Group, b Group) bool {
				if a.Cards.Find(CardIs(hand.Last)).Size() > 0 && b.Cards.Find(CardIs(hand.Last)).Size() > 0 {
					aMax := a.Cards.FindMaxPoint()
					bMax := b.Cards.FindMaxPoint()
					if Abs(aMax.Point-bMax.Point) == 1 {
						return true
					}
				}
				return false
			})
		},
	}

	DAN_DIAO_JIANG = Fan{
		Name: "单调将",
		Fan:  1,
		Match: func(hand GroupHand) bool {

			return hand.FindLastGroup().HasGroup(func(g Group) bool {
				return g.isJiang()
			})
		},
	}

	ZI_MO = Fan{
		Name: "自摸",
		Fan:  1,
		Match: func(hand GroupHand) bool {
			return hand.ZiMo
		},
	}
)
