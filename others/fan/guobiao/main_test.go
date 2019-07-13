package guobiao

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcFan(t *testing.T) {
	hand, err := Parse("碰东东东, 万234 筒234 条45677 , 摸条5")
	assert.NoError(t, err)
	fan := CalcFan(hand)
	fmt.Println(fan)
}
