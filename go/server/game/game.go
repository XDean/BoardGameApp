package game

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
