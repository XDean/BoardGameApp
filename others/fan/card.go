package fan

import "fmt"

const (
	// type
	TIAO CardType = iota
	TONG
	WAN
	ZI
	HUA
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

const (
	// hua
	H_CHUN int = iota + 1
	H_XIA
	H_QIU
	H_DONG
	H_MEI
	H_LAN
	H_ZHU
	H_JU
)

type (
	CardType int
	Card     struct {
		Type  CardType
		Point int
	}
)

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
	case HUA:
		return c.Point > 0 && c.Point < 9
	default:
		return false
	}
}

func (c Card) String() string {
	if !c.isValid() {
		return "[无效牌]"
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
	case HUA:
		switch c.Point {
		case H_CHUN:
			return "春"
		case H_XIA:
			return "夏"
		case H_QIU:
			return "秋"
		case H_DONG:
			return "冬"
		case H_MEI:
			return "梅"
		case H_LAN:
			return "兰"
		case H_ZHU:
			return "竹"
		case H_JU:
			return "菊"
		}
	}
	panic("never happen")
}
