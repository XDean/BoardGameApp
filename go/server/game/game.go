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

var Games []*Game

func Register(game *Game) {
	Games = append(Games, game)
}

func FindGame(id string) (*Game, error) {
	for _, v := range Games {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, errors.New("No such game")
}
