package fan

type (
	Hu func(hand Hand) []GroupHand
)

var HU_COMMON = func(hand Hand) []GroupHand {
	result := make([]GroupHand, 0)

	for card, _ := range hand.Private {
		for _, t := range []GroupType{SHI_SAN_YAO, QI_XING_BU_KAO, QUAN_BU_KAO, QI_DUI} {
			if ok, group, _ := t.Find(hand.Private, card); ok {
				return []GroupHand{
					{
						Groups: []Group{group},
						Last:   hand.Last,
						ZiMo:   hand.ZiMo,
					},
				}
			}
		}

		break
	}

	return result
}
