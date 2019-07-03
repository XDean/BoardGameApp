package main

const (
	F_COMMON FanType = iota
	F_HU
	F_EXTRA
)

type (
	FanType int
	Fan     struct {
		Name   string
		Fan    int
		Type   FanType
		Ignore []string
		Match  func(hand Hand) bool
	}
)

var (
	DA_SAN_YUAN = Fan{
		Name: "大三元",
		Fan:  88,
		Match: func(hand Hand) bool {
			return true
		},
	}
)
