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
			Groups: append(hand.Public, groups...),
			Cards:  cards,
			Last:   hand.Last,
			ZiMo:   hand.ZiMo,
		})
	}
	return result
}

func findGroup(cards Cards) [][]Group {
	if cards.Size() == 0 {
		return [][]Group{[]Group{}}
	}
	result := make([][]Group, 0)
	for card, count := range cards {
		if count == 0 {
			continue
		}
		for _, t := range []GroupType{GT_SHI_SAN_YAO, GT_QI_XING_BU_KAO, GT_QUAN_BU_KAO, GT_QI_DUI} {
			if ok, group, _ := t.Find(cards, card); ok {
				result = append(result, []Group{group})
			}
		}
		if cards.Size()%3 == 2 {
			if ok, group, left := GT_JIANG.Find(cards, card); ok {
				leftGroups := findGroup(left)
				for i, g := range leftGroups {
					leftGroups[i] = append(g, group)
				}
				result = append(result, leftGroups...)
			}
		}
		for _, t := range []GroupType{GT_ZU_HE_LONG, GT_KE, GT_SHUN} {
			if ok, group, left := t.Find(cards, card); ok {
				leftGroups := findGroup(left)
				for i, g := range leftGroups {
					leftGroups[i] = append(g, group)
				}
				result = append(result, leftGroups...)
			}
		}
	}
	return result
}
