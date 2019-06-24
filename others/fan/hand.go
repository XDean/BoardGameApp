package fan

import (
	"errors"
)

const (
	CHI GroupType = iota
	PENG
	MING_GANG
	AN_GANG
)

type (
	GroupType int
	Group     struct {
		Type  GroupType
		Start Card
	}

	Hand struct {
		Public  []Group
		Private []Card
		Last    Card
		ZiMo    bool
	}
)

func (h Hand) isValid() (bool, error) {
	if len(h.Private)%3 != 1 {
		return false, errors.New("牌数应为13张")
	}
	for _, g := range h.Public {
		if ok, err := g.isValid(); !ok {
			return false, err
		}
	}
	for _, c := range h.Private {
		if c.isValid() {
			return false, errors.New(c.String())
		}
	}
	cards := h.AllCards()
	cardCount := make(map[Card]int)
	for _, card := range cards {
		if count, ok := cardCount[card]; ok {
			if count == 4 {
				return false, errors.New("不存在超过四张同样的牌： " + card.String())
			}
		}
	}
	return true, nil
}

func (g Group) isValid() (bool, error) {
	if !g.Start.isValid() {
		return false, errors.New("无效组合")
	}
	if g.Type == CHI {
		if !g.Start.isTTW() {
			return false, errors.New("无效组合，字不能吃")
		}
		if g.Start.Point > 7 {
			return false, errors.New("顺子不能超过9")
		}
	}
	return true, nil
}

func (g Group) AllCards() []Card {
	cards := make([]Card, 0)
	cards = append(cards, g.Start)
	switch g.Type {
	case CHI:
		cards = append(cards, g.Start.NextPoint())
	case AN_GANG:
		fallthrough
	case MING_GANG:
		cards = append(cards, g.Start.Copy())
		fallthrough
	case PENG:
		cards = append(cards, g.Start.Copy())
		cards = append(cards, g.Start.Copy())
	}
	return cards
}

func (h Hand) AllCards() []Card {
	cards := make([]Card, 0)
	for _, g := range h.Public {
		cards = append(cards, g.AllCards()...)
	}
	cards = append(cards, h.Private...)
	cards = append(cards, h.Last)
	return cards
}
