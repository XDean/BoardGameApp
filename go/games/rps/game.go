package rps

import (
	"errors"
	"github.com/xdean/miniboardgame/go/server/game"
	"github.com/xdean/miniboardgame/go/server/model/space"
)

type (
	Value int
	Game  struct {
		Count  int
		Values []Value
		Lose   []bool
		Winner int
		Thread space.Publisher
	}
	Event struct {
		game.BaseEvent
		Value Value
	}
)

const (
	ROCK     = 0x100
	PAPER    = 0x010
	SCISSORS = 0x001
)

var (
	ALL = []Value{ROCK, PAPER, SCISSORS}
)

func NewGame(i int) *Game {
	return &Game{
		Count:  i,
		Values: make([]Value, i),
		Lose:   make([]bool, i),
	}
}

func (g *Game) Play(seat int, value Value) game.Response {
	if g.Lose[seat] {
		return errors.New("You had lost")
	}
	if g.Values[seat] != 0 {
		return errors.New("You had given your choice")
	}
	g.Values[seat] = value
	defer func() {
		for i, v := range g.Values {
			if !g.Lose[i] && v == 0 {
				return
			}
		}
		g.Thread.SendEvent(space.Message{
			To:      -1,
			Topic:   "Win",
			Payload: g.winner(),
		})
	}()
	return "Accept"
}

func (g *Game) winner() []int {
	mask := 0x0
	for i, v := range g.Values {
		if !(g.Lose[i]) {
			mask |= int(v)
		}
	}
	var win Value = 0
	if mask == ROCK+PAPER {
		win = PAPER
	} else if mask == ROCK+SCISSORS {
		win = ROCK
	} else if mask == PAPER+SCISSORS {
		win = SCISSORS
	}
	if win == 0 {
		for i := range g.Values {
			g.Values[i] = 0
		}
		return []int{}
	}
	winner := make([]int, 0)
	for i, v := range g.Values {
		if !(g.Lose[i]) {
			if v == win {
				winner = append(winner, i)
				g.Values[i] = 0
			} else {
				g.Lose[i] = true
			}
		}
	}
	if len(winner) == 1 {
		g.Winner = winner[0]
	}
	return winner
}
