package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckSignature(t *testing.T) {
	assert.True(t, checkSignature("test", "bb5f8aca3af358288f3fb866b7580e9378be123d", "1522033654", "651288971"))
	assert.False(t, checkSignature("no", "bb5f8aca3af358288f3fb866b7580e9378be123d", "1522033654", "651288971"))

	assert.True(t, checkSignature("XDeanMiniBoardgame", "069a8f61dc6ab60454c70d36713768a38680f0c0", "1563451317", "127831151"))
}
