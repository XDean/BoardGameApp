package guobiao

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcFan(t *testing.T) {
	testFan(t, "碰万2, 万234 筒222 条45677 , 摸7",
		DUAN_YAO, SI_GUI_YI, SHUANG_TONG_KE, WU_ZI, BIAN_ZHANG, DAN_DIAO_JIANG, ZI_MO)
	testFan(t, "暗东 暗西 暗南 暗北，中中，摸中",
		DA_SI_XI, SI_GANG, ZI_YI_SE, SI_AN_KE, BU_QIU_REN)
}

func testFan(t *testing.T, handStr string, expected ...Fan) {
	hand, err := Parse(handStr)
	assert.NoError(t, err)
	actual := CalcFan(hand)
	assert.Equal(t, len(expected), len(actual))
	fail := len(expected) == len(actual)
	for _, e := range expected {
		found := false
		for _, a := range actual {
			if a.Name == e.Name {
				found = true
				break
			}
		}
		if !found {
			fail = true
			assert.Fail(t, "Expect %s, But %s", expected, actual)
			break
		}
	}
	if !fail {
		t.Log("["+handStr+"]", "胡", actual, "共计", actual.Sum(), "番")
	}
}
