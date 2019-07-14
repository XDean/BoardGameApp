package guobiao

import "sort"

func isEqual(a, b, c interface{}) bool {
	return a == b && b == c
}

func isTBW(a, b, c CardType) bool {
	return a != ZI && b != ZI && c != ZI && a != b && a != c && b != c
}

func isStep(a, b, c int) bool {
	s := []int{a, b, c}
	sort.Ints(s)
	return s[1]-s[0] == 1 && s[2]-s[1] == 1
}

func isStep3(a, b, c int) bool {
	s := []int{a, b, c}
	sort.Ints(s)
	return s[1]-s[0] == 3 && s[2]-s[1] == 3
}
