package guobiao

import "sort"

func isEqual(is ...interface{}) bool {
	for i := 1; i < len(is); i++ {
		if is[i] != is[i-1] {
			return false
		}
	}
	return true
}

func isTBW(a, b, c CardType) bool {
	return a != ZI && b != ZI && c != ZI && a != b && a != c && b != c
}

func isStep(is ...int) bool {
	sort.Ints(is)
	for i := 1; i < len(is); i++ {
		if is[i]-is[i-1] != 1 {
			return false
		}
	}
	return true
}

func isStep3(is ...int) bool {
	sort.Ints(is)
	for i := 1; i < len(is); i++ {
		if is[i]-is[i-1] != 3 {
			return false
		}
	}
	return true
}
