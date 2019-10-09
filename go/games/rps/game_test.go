package rps

import (
	"github.com/stretchr/testify/assert"
	"github.com/xdean/miniboardgame/go/server/game"
	"github.com/xdean/miniboardgame/go/server/model/space"
	"testing"
)

type Receiver int

const host Receiver = 1

func (Receiver) EventHostId() string {
	return "test"
}

func TestGame_Play(t *testing.T) {
	g := NewGame(5)
	g.Thread = space.Publish(host)

	noErr(t, g.Play(0, ROCK))
	mustErr(t, g.Play(0, ROCK))
	noErr(t, g.Play(1, ROCK))
	noErr(t, g.Play(2, ROCK))
	noErr(t, g.Play(3, ROCK))
	noErr(t, g.Play(4, ROCK))

	noErr(t, g.Play(0, ROCK))
	noErr(t, g.Play(1, ROCK))
	noErr(t, g.Play(2, ROCK))
	noErr(t, g.Play(3, ROCK))
	noErr(t, g.Play(4, SCISSORS))

	noErr(t, g.Play(0, PAPER))
	noErr(t, g.Play(1, ROCK))
	noErr(t, g.Play(2, PAPER))
	noErr(t, g.Play(3, ROCK))
	mustErr(t, g.Play(4, SCISSORS))

	noErr(t, g.Play(0, PAPER))
	noErr(t, g.Play(2, SCISSORS))

	assert.Equal(t, 2, g.Winner)
}

func mustErr(t *testing.T, res game.Response) {
	if _, ok := res.(error); !ok {
		t.Errorf("Except error but: %s", res)
	}
}

func noErr(t *testing.T, res game.Response) {
	if err, ok := res.(error); ok {
		assert.NoError(t, err)
	}
}
