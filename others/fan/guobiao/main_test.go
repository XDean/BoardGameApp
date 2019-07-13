package guobiao

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcFan(t *testing.T) {
	hand, err := Parse("吃筒123 明中中中中, 万33 筒234 北北北, 摸北")
	assert.NoError(t, err)
	fan := CalcFan(hand)
	fmt.Println(fan)
}
