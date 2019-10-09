package game

import "github.com/pkg/errors"

type (
	Game struct {
		Id   string
		Name string

		// setup
		Player  Range
		Options []Option

		// play
		NewEvent func() Event
		OnEvent  func(Event) Response
	}
)

var Games = map[string]*Game{}

func Register(game *Game) {
	Games[game.Id] = game
}

func FindGame(id string) (*Game, error) {
	if game, ok := Games[id]; ok {
		return game, nil
	} else {
		return nil, errors.New("No such game")
	}
}
