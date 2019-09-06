package guobiao

var (
	ALL_48_FAN = []Fan{
		YI_SE_SI_TONG_SHUN,
		YI_SE_SI_JIE_GAO,
	}
	YI_SE_SI_TONG_SHUN = Fan{
		Name:   "一色四同顺",
		Fan:    48,
		Ignore: []string{YI_SE_SAN_TONG_SHUN.Name, YI_BAN_GAO.Name, SI_GUI_YI.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasQuad(func(a, b, c, d Group) bool {
				if a.isShunZi() && b.isShunZi() && c.isShunZi() && d.isShunZi() {
					am := a.Cards.FindMinPointCard()
					bm := b.Cards.FindMinPointCard()
					cm := c.Cards.FindMinPointCard()
					dm := d.Cards.FindMinPointCard()
					return isEqual(am, bm, cm, dm)
				}
				return false
			})
		},
	}
	YI_SE_SI_JIE_GAO = Fan{
		Name:   "一色四节高",
		Fan:    48,
		Ignore: []string{YI_SE_SAN_JIE_GAO.Name, PENG_PENG_HU.Name},
		Match: func(hand GroupHand) bool {
			return hand.Groups.HasQuad(func(a, b, c, d Group) bool {
				if a.isKeZi() && b.isKeZi() && c.isKeZi() && d.isKeZi() {
					am := a.Cards.Any()
					bm := b.Cards.Any()
					cm := c.Cards.Any()
					dm := d.Cards.Any()
					if isEqual(am.Type, bm.Type, cm.Type, dm.Type) && am.Type != ZI {
						return isStep(am.Point, bm.Point, cm.Point, dm.Point)
					}
				}
				return false
			})
		},
	}
)
