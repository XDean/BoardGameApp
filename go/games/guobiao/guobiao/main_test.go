package guobiao

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcFan(t *testing.T) {
	hand, err := Parse("碰万222, 万234 筒222 条45677 , 摸7")
	assert.NoError(t, err)
	fan := CalcFan(hand)
	fmt.Println(fan, fan.Sum())
}
