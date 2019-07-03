package main

import (
	"errors"
)

const (
	CHI GroupType = iota
	PENG
	MING_GANG
	AN_GANG

	KE
	SHUN
	DUI
	// special type
	ZU_HE_LONG
	QUAN_BU_KAO
	QI_XING_BU_KAO
	SHI_SAN_YAO
)

var (
	ALL_PRIVATE_GROUP_TYPE = []GroupType{
		KE, SHUN, DUI, ZU_HE_LONG, QUAN_BU_KAO, QI_XING_BU_KAO, SHI_SAN_YAO,
	}
)

type (
	Cards       map[Card]int
	GroupType   int
	PublicType  GroupType
	PrivateType GroupType

	Group struct {
		Type  GroupType
		Start Card
		Cards Cards // for special type
	}

	Hand struct {
		Public  []Group
		Private Cards
		Last    Card
		ZiMo    bool
	}

	GroupedHand struct {
		Public  []Group
		Private []Group
		Last    Card
		ZiMo    bool
	}
)

func (t GroupType) IsPublic() bool {
	return t == PENG || t == CHI || t == AN_GANG || t == MING_GANG
}

func (t GroupType) IsSpecial() bool {
	return t == ZU_HE_LONG || t == QUAN_BU_KAO || t == QI_XING_BU_KAO || t == SHI_SAN_YAO
}

func (h Hand) isValid() (bool, error) {
	if len(h.Private)%3 != 1 {
		return false, errors.New("牌数应为13张")
	}
	for _, g := range h.Public {
		if !g.Type.IsPublic() {
			return false, errors.New("只支持碰吃杠")
		}
		if ok, err := g.isValid(); !ok {
			return false, err
		}
	}
	for k, _ := range h.Private {
		if !k.isValid() {
			return false, errors.New(k.String())
		}
	}
	cards := h.AllCards()
	cardCount := make(Cards)
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
	if g.Type == GroupType(CHI) {
		if !g.Start.isTBW() {
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
	for k, v := range h.Private {
		for i := 0; i < v; i++ {
			cards = append(cards, k)
		}
	}
	cards = append(cards, h.Last)
	return cards
}

func (h GroupedHand) Copy() GroupedHand {
	publicGroup := make([]Group, len(h.Public))
	privateGroup := make([]Group, len(h.Private))
	copy(publicGroup, h.Public)
	copy(privateGroup, h.Private)
	h.Public = publicGroup
	h.Private = privateGroup
	return h
}

func (c Cards) IsValid() bool {
	for _, v := range c {
		if v < 0 {
			return false
		}
	}
	return true
}

func (c Cards) Size() int {
	i := 0
	for _, v := range c {
		i += v
	}
	return i
}

func (c Cards) Copy() Cards {
	cards := make(Cards)
	for k, v := range c {
		cards[k] = v
	}
	return cards
}
