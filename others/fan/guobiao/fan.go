package guobiao

type (
	Fan struct {
		Name   string
		Fan    int
		Ignore []string
		Match  func(hand GroupHand) bool
	}

	ExtraFan struct {
		Name   string
		Fan    int
		Ignore []string
	}
)

func CalcFan(hand GroupHand) []Fan {
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

var (
	ALL_FAN []Fan
)
