package guobiao

var (
	ALL_64_FAN = []Fan{
		XIAO_SI_XI,
		XIAO_SAN_YUAN,
		ZI_YI_SE,
		YI_SE_SHUANG_LONG_HUI,
		QING_YAO_JIU,
		SI_AN_KE,
	}
	XIAO_SI_XI = Fan{
		Name:   "小四喜",
		Fan:    64,
		Ignore: []string{SAN_FENG_KE.Name, YAO_JIU_KE.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Has(func(g Group) bool {
				return g.isJiang() && g.Cards.Any().isFeng()
			}) && hand.Groups.Find(func(g Group) bool {
				return g.isKeZi() && g.Cards.Any().isFeng()
			}).Size() == 3
		},
	}
	XIAO_SAN_YUAN = Fan{
		Name:   "小三元",
		Fan:    64,
		Ignore: []string{JIAN_KE.Name, SHUANG_JIAN_KE.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Has(func(g Group) bool {
				return g.isJiang() && g.Cards.Any().isZFB()
			}) && hand.Groups.Find(func(g Group) bool {
				return g.isKeZi() && g.Cards.Any().isZFB()
			}).Size() == 2
		},
	}
	ZI_YI_SE = Fan{
		Name:   "字一色",
		Fan:    64,
		Ignore: []string{PENG_PENG_HU.Name, QUAN_DAI_YAO.Name, YAO_JIU_KE.Name, QUE_YI_MEN.Name},
		Match: func(hand GroupHand) bool {
			return hand.Cards.All(TypeIs(ZI))
		},
	}
	YI_SE_SHUANG_LONG_HUI = Fan{
		Name:   "一色双龙会",
		Fan:    64,
		Ignore: []string{QING_YI_SE.Name, YI_BAN_GAO.Name, LAO_SHAO_FU.Name, WU_ZI.Name, PING_HU.Name},
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
			if jiang.Size() != 1 || left.Size() != 4 {
				return false
			}
			return isEqual(jiang[0].Cards.Any().Type, left[0].Cards.Any().Type)
		},
	}
	QING_YAO_JIU = Fan{
		Name:   "清幺九",
		Fan:    64,
		Ignore: []string{PENG_PENG_HU.Name, QUAN_DAI_YAO.Name, YAO_JIU_KE.Name, WU_ZI.Name},
		Match: func(hand GroupHand) bool {
			return hand.Cards.All(PointIs(1).Or(PointIs(9)))
		},
	}
	SI_AN_KE = Fan{
		Name:   "四暗刻",
		Fan:    64,
		Ignore: []string{SAN_AN_KE.Name, PENG_PENG_HU.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.Find(func(g Group) bool {
				return g.isAnKe()
			}).Size() == 4
		},
	}
)
