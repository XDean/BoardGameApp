package guobiao

var (
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
			return !hand.Groups.HasGroupPair(func(a Group, b Group) bool {
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
)
