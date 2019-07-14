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
		if ignore[f.Name] {
			continue
		}
		if f.Match(hand) {
			result = append(result, f)
			for _, i := range f.Ignore {
				ignore[i] = true
			}
		}
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
	ALL_FAN = append(
		ALL_2_FAN,
		ALL_1_FAN...)
)
