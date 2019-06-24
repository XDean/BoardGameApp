package fan

import "fmt"

const (
	// type
	TIAO CardType = iota
	TONG
	WAN
	ZI
)
const (
	// zi
	Z_DONG int = iota + 1
	Z_NAN
	Z_XI
	Z_BEI
	Z_ZHONG
	Z_FA
	Z_BAI
)

type (
	CardType int
	Card     struct {
		Type  CardType
		Point int
	}
)

func (c Card) isTTW() bool {
	return c.Type == TIAO || c.Type == TONG || c.Type == WAN
}

func (c Card) isZi() bool {
	return c.Type == ZI
}

func (c Card) Copy() Card {
	return c
}

func (c Card) NextPoint() Card {
	c.Point++
	return c
}

func (c Card) isValid() bool {
	switch c.Type {
	case TIAO:
		fallthrough
	case TONG:
		fallthrough
	case WAN:
		return c.Point > 0 && c.Point < 10
	case ZI:
		return c.Point > 0 && c.Point < 8
	default:
		return false
	}
}

func (c Card) String() string {
	if !c.isValid() {
		return fmt.Sprintf("[无效牌 %d %d]", c.Type, c.Point)
	}
	switch c.Type {
	case TIAO:
		return fmt.Sprintf("%d条", c.Point)
	case TONG:
		return fmt.Sprintf("%d筒", c.Point)
	case WAN:
		return fmt.Sprintf("%d万", c.Point)
	case ZI:
		switch c.Point {
		case Z_DONG:
			return "东"
		case Z_NAN:
			return "南"
		case Z_XI:
			return "西"
		case Z_BEI:
			return "北"
		case Z_ZHONG:
			return "中"
		case Z_FA:
			return "发"
		case Z_BAI:
			return "白"
		}
	}
	panic("never happen")
}
