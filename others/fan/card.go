package fan

import "fmt"

const (
	// type
	TIAO int = iota
	TONG
	WAN
	ZI
	HUA

	// zi
	Z_DONG int = iota
	Z_NAN
	Z_XI
	Z_BEI
	Z_ZHONG
	Z_FA
	Z_BAI

	// hua
	H_CHUN int = iota
	H_XIA
	H_QIU
	H_DONG
	H_MEI
	H_LAN
	H_ZHU
	H_JU
)

type (
	Point int
	Card  struct {
		Type  int
		Point int
	}
)

func (c Card) String() string {
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
	return "[无效牌]"
}
