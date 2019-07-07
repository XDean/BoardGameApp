package guobiao

func FindHu(hand Hand) []GroupHand {
	result := make([]GroupHand, 0)
	cards := Cards{}
	for _, g := range hand.Public {
		for card, count := range g.Cards {
			cards[card] += count
		}
	}
	for card, count := range hand.Private {
		cards[card] += count
	}
	for _, groups := range findGroup(hand.Private) {
		result = append(result, GroupHand{
			Groups: groups,
			Cards:  cards,
			Last:   hand.Last,
			ZiMo:   hand.ZiMo,
		})
	}
	return result
}

func findGroup(cards Cards) [][]Group {
	for card, _ := range cards {
		for _, t := range []GroupType{GT_SHI_SAN_YAO, GT_QI_XING_BU_KAO, GT_QUAN_BU_KAO, GT_QI_DUI} {
			if ok, group, _ := t.Find(cards, card); ok {
				return [][]Group{{group}}
			}
		}
		for _, t := range []GroupType{GT_ZU_HE_LONG, GT_KE, GT_SHUN} {
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
