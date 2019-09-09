package game

type (
	Game struct {
		Id   string
		Name string

		// setup
		Player  Range
		Options []Option

		// play
		NewEvent func() interface{}
		OnEvent  func(interface{}) (interface{}, error)
	}

	Option struct {
		Id     string
		Name   string
		Type   OptionType
		Domain interface{}
	}

	OptionType int

	Range struct {
		Min int
		Max int
	}

	Event struct {
		PlayerId int
		RoomId   int
	}
)

var Games []*Game

func Register(game *Game) {
	Games = append(Games, game)
}
