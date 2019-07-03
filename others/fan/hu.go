package main

type (
	Hu struct {
		Name   string
		Fan    int
		Ignore []Fan
		Group  func(hand Hand) []GroupedHand
	}
)

var (
	HU_COMMON = Hu{
		Name:  "",
		Fan:   0,
		Group: toCommonGroup,
	}
)

func toCommonGroup(hand Hand) []GroupedHand {
	origin := GroupedHand{
		Public:  hand.Public,
		Private: []Group{},
		Last:    hand.Last,
		ZiMo:    hand.ZiMo,
	}
	return findCommonGroup(origin, hand.Private)
}

func findCommonGroup(hand GroupedHand, cards Cards) []GroupedHand {
	result := make([]GroupedHand, 0)
	size := cards.Size()
	if size == 0 {
		return nil
	} else {
		for card, count := range cards {
			for _, groupType := range ALL_PRIVATE_GROUP_TYPE {
				if ok, group, left := cards.Find(groupType, card); ok {

				}
			}
		}
	}
}

func (c Cards) Find(t GroupType, card Card) (ok bool, group Group, left Cards) {
	size := c.Size()
	ok = false
	group = Group{
		Type:  t,
		Start: card,
	}
	left = c.Copy()
	switch t {
	case KE:
		left[card] -= 3
	case SHUN:
		next := card.NextPoint()
		next2 := next.NextPoint()
		left[card] -= 1
		left[next] -= 1
		left[next2] -= 1
	case DUI:
		left[card] -= 2
	case ZU_HE_LONG:
		if size >= 9 {
		zhl:
			for tiao := 1; tiao <= 3; tiao++ {
				for bing := 1; bing <= 3; bing++ {
					for wan := 1; wan <= 3; wan++ {
						if tiao != bing && tiao != wan && bing != wan {
							t := Card{Type: TIAO, Point: tiao}
							b := Card{Type: BING, Point: bing}
							w := Card{Type: WAN, Point: wan}
							copy := left.Copy()
							copy[t] -= 1
							copy[t.Next(3)] -= 1
							copy[t.Next(6)] -= 1
							copy[b] -= 1
							copy[b.Next(3)] -= 1
							copy[b.Next(6)] -= 1
							copy[w] -= 1
							copy[w.Next(3)] -= 1
							copy[w.Next(6)] -= 1
							if copy.IsValid() {
								left = copy
								break zhl
							}
						}
					}
				}
			}
		}
	case QUAN_BU_KAO:
	case QI_XING_BU_KAO:
	case SHI_SAN_YAO:
	}
	ok = left.IsValid()
	return
}
