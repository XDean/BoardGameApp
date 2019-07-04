package fan

func FindHu(hand Hand) []GroupHand {
	result := make([]GroupHand, 0)
	for _, groups := range findGroup(hand.Private) {
		result = append(result, GroupHand{
			Groups: groups,
			Last:   hand.Last,
			ZiMo:   hand.ZiMo,
		})
	}
	return result
}

func findGroup(cards Cards) [][]Group {
	for card, _ := range cards {
		for _, t := range []GroupType{SHI_SAN_YAO, QI_XING_BU_KAO, QUAN_BU_KAO, QI_DUI} {
			if ok, group, _ := t.Find(cards, card); ok {
				return [][]Group{{group}}
			}
		}
		for _, t := range []GroupType{ZU_HE_LONG, KE, SHUN} {
			if ok, group, left := t.Find(cards, card); ok {
				leftGroups := findGroup(left)
				for i, g := range leftGroups {
					leftGroups[i] = append(g, group)
				}
				return leftGroups
			}
		}

		break
	}
	return nil
}
