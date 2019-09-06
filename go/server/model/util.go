package model

import "math"

type RowBound struct {
	Limit  int
	Offset int
}

var Unbound = RowBound{math.MaxInt32, 0}
