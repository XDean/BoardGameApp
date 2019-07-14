package guobiao

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	parseAndFormat(t,
		"吃筒123 明中中中中, 万33 筒234 北北北, 摸北",
		"吃饼123 明中中中中 , 饼234 万33 北北北 , 摸北")
}

func parseAndFormat(t *testing.T, origin, format string) {
	hand, err := Parse(origin)
	assert.NoError(t, err)

	result := Format(hand)
	assert.Equal(t, format, result)

	hand2, err := Parse(result)
	assert.NoError(t, err)
	assert.Equal(t, hand, hand2)

	result2 := Format(hand)
	assert.Equal(t, format, result2)

}
