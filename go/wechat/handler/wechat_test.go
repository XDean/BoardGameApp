package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckSignature(t *testing.T) {
	assert.True(t, checkSignature("test", "bb5f8aca3af358288f3fb866b7580e9378be123d", "1522033654", "651288971"))
	assert.False(t, checkSignature("no", "bb5f8aca3af358288f3fb866b7580e9378be123d", "1522033654", "651288971"))
}
