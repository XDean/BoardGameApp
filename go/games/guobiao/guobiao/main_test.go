package guobiao

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcFan(t *testing.T) {
	testFan(t, "碰万222, 万234 筒222 条45677 , 摸7", DUAN_YAO, SI_GUI_YI, SHUANG_TONG_KE, WU_ZI, BIAN_ZHANG, DAN_DIAO_JIANG, ZI_MO)
}

func testFan(t *testing.T, handStr string, expected ...Fan) {
	hand, err := Parse(handStr)
	assert.NoError(t, err)
	actual := CalcFan(hand)
	assert.Equal(t, len(actual), len(expected))
	for _, e := range expected {
		found := false
		for _, a := range actual {
			if a.Name == e.Name {
				found = true
				break
			}
		}
		if !found {
			assert.Fail(t, "Expect %s", e)
		}
	}
	t.Log("["+handStr+"]", "胡", actual, "共计", actual.Sum(), "番")
}
