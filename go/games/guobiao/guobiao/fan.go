package guobiao

import "fmt"

type (
	Fan struct {
		Name   string
		Fan    int
		Ignore []string
		Match  func(hand GroupHand) bool
	}
	Fans []Fan

	ExtraFan struct {
		Name   string
		Fan    int
		Ignore []string
	}
)

func (f Fan) String() string {
	return fmt.Sprintf("%s(%d)", f.Name, f.Fan)
}

func CalcGroupFan(hand GroupHand) Fans {
	result := make([]Fan, 0)
	ignore := make(map[string]bool, 0)
	for _, f := range ALL_FAN {
		if f.Match(hand) {
			for _, i := range f.Ignore {
				ignore[i] = true
			}
			if !ignore[f.Name] {
				result = append(result, f)
			}
		}
	}
	if len(result) == 0 {
		result = append(result, WU_FAN_HU)
	}
	return result
}

func (f Fans) Sum() int {
	sum := 0
	for _, fan := range f {
		sum += fan.Fan
	}
	return sum
}

var (
	ALL_FAN = appendFan(
		ALL_88_FAN,
		ALL_64_FAN,
		ALL_48_FAN,
		ALL_32_FAN,
		ALL_24_FAN,
		ALL_16_FAN,
		ALL_12_FAN,
		ALL_8_FAN,
		ALL_6_FAN,
		ALL_4_FAN,
		ALL_2_FAN,
		ALL_1_FAN,
	)
)

func appendFan(fans ...[]Fan) []Fan {
	result := make([]Fan, 0)
	for _, v := range fans {
		result = append(result, v...)
	}
	return result
}
