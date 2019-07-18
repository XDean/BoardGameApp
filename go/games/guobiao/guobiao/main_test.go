package guobiao

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcFan(t *testing.T) {
	testFan(t, "碰万2, 万234 筒222 条45677 , 摸7",
		DUAN_YAO, SI_GUI_YI, SHUANG_TONG_KE, WU_ZI, BIAN_ZHANG, DAN_DIAO_JIANG, ZI_MO)
	testFan(t, "暗东 暗西 暗南 暗北，中中，摸中",
		DA_SI_XI, SI_GANG, ZI_YI_SE, SI_AN_KE, BU_QIU_REN)
	testFan(t, "，条11122345678999，摸2",
		JIU_LIAN_BAO_DENG, BU_QIU_REN, SHUANG_AN_KE, LIAN_LIU, DAN_DIAO_JIANG)
	testFan(t, "，条11223355778899，3",
		YI_SE_SHUANG_LONG_HUI, MEN_QIAN_QING, BIAN_ZHANG)
	testFan(t, "吃条123，条12355778899，3",
		YI_SE_SHUANG_LONG_HUI, BIAN_ZHANG)
	testFan(t, "条碰2 碰3 碰8, 23444，4",
		LV_YI_SE, QING_YI_SE, DUAN_YAO, SI_GUI_YI, DAN_DIAO_JIANG)
	testFan(t, "，东东南南西西北北中中发发白白，白",
		ZI_YI_SE, QI_DUI)
	testFan(t, "，中中中发发发白白白万55566，摸5",
		DA_SAN_YUAN, SI_AN_KE, HUN_YI_SE, BU_QIU_REN)
	testFan(t, "碰万1，22233344455，5",
		YI_SE_SI_JIE_GAO, QING_YI_SE, SAN_AN_KE, YAO_JIU_KE, DAN_DIAO_JIANG)
}

func testFan(t *testing.T, handStr string, expected ...Fan) {
	hand, err := Parse(handStr)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	actual := CalcFan(hand)
	assert.Equal(t, len(expected), len(actual))
	for _, e := range expected {
		found := false
		for _, a := range actual {
			if a.Name == e.Name {
				found = true
				break
			}
		}
		if !found {
			assert.Fail(t, fmt.Sprintf("Expect %v, But %v", expected, actual))
			break
		}
	}
	t.Log("["+handStr+"]", "胡", actual, "共计", actual.Sum(), "番")
}
