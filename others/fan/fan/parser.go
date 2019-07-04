package fan

import (
	"fmt"
	"github.com/pkg/errors"
)

// Ct123 b456, DDDZZw555, w5

const (
	mode_public = iota
	mode_private
	mode_last
)

func Parse(str string) (Hand, error) {
	cardType := ZI

	mode := mode_public

	publicCards := Cards{}
	publicType := QI_XING_BU_KAO

	newPublicType := QI_XING_BU_KAO
	newCard := NIL_CARD

	hand := Hand{
		Public:  []Group{},
		Private: Cards{},
		Last:    NIL_CARD,
	}

	for i, c := range str {
		switch c {
		case ' ', '\t':
			continue
		case ',':
			switch mode {
			case mode_public:
				if publicType.Public || publicCards.Size() != 0 {
					return hand, parseError(i, "未完成的吃碰杠")
				}
				mode = mode_private
			case mode_private:
				mode = mode_last
			case mode_last:
				return hand, parseError(i, "多余的块")
			}
		case '条':
			cardType = TIAO
		case '饼':
			cardType = BING
		case '万':
			cardType = WAN
		case '吃':
			newPublicType = CHI
		case '碰':
			newPublicType = PENG
		case '暗':
			newPublicType = AN_GANG
		case '明':
			newPublicType = MING_GANG
		case '摸':
			if hand.ZiMo {
				return hand, parseError(i, "冗余的自摸")
			}
			hand.ZiMo = true
		case '东':
			newCard = DONG_CARD
		case '南':
			newCard = NAN_CARD
		case '西':
			newCard = XI_CARD
		case '北':
			newCard = BEI_CARD
		case '中':
			newCard = ZHONG_CARD
		case '发':
			newCard = FA_CARD
		case '白':
			newCard = BAI_CARD
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if cardType == ZI {
				return hand, parseError(i, "未指定牌类型")
			}
			newCard = Card{Type: cardType, Point: int(c)}
		}
		if newPublicType.Public {
			if mode != mode_public {
				return hand, parseError(i, "只能在第一部分吃碰杠")
			}
			if publicType.Public {
				return hand, parseError(i, "重复吃碰杠")
			}
			publicType = newPublicType
			newPublicType = QI_XING_BU_KAO
		}
		if newCard != NIL_CARD {
			switch mode {
			case mode_public:
				if publicType.Public {
					publicCards[newCard] += 1
					if publicCards.Size() == publicType.CardCount {
						if ok, group, _ := publicType.Find(publicCards, newCard); ok {
							hand.Public = append(hand.Public, group)
							publicType = QI_XING_BU_KAO
							publicCards = Cards{}
						} else {
							return hand, parseError(i, "错误的牌组合")
						}
					}
				} else {
					return hand, parseError(i, "未指定吃碰杠类型")
				}
			case mode_private:
				hand.Private[newCard] += 1
			case mode_last:
				if hand.Last == NIL_CARD {
					hand.Last = newCard
				} else {
					return hand, parseError(i, "多余的胡牌")
				}
			}
			newCard = NIL_CARD
		}
	}
	if mode != mode_last {
		return hand, parseError(len(str), "缺少胡牌块")
	}
	if hand.Last == NIL_CARD {
		return hand, parseError(len(str), "缺少胡牌")
	}
	if hand.CardCount() != 14 {
		return hand, parseError(len(str), "牌数应为14张")
	}
	return hand, nil
}

func parseError(index int, message string) error {
	return errors.New(fmt.Sprintf("Parse failed at {%d} because: %s", index, message))
}
