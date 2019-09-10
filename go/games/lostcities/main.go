package lostcities

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/xdean/miniboardgame/go/server/game"
)

var Instance = &game.Game{
	Id:   "lost-cities",
	Name: "失落的城市",
	Player: game.Range{
		Min: 2,
		Max: 2,
	},
	Options: nil,
	NewEvent: func() game.Event {
		return Event{}
	},
	OnEvent: func(e game.Event) game.Response { // card or error
		eventStream <- e
		return e.GetResponse()
	},
}

var eventStream = make(chan game.Event)
var games = make(map[int]*Game)

func init() {
	go run()
}

func run() {
	for {
		select {
		case event := <-eventStream:
			logrus.Debug(event)
			switch t := event.(type) {
			case Event:
				g := games[t.GetRoomId()]
				t.PutResponse(g.Play(t))
			case game.NewGameEvent:
				g := NewStandardGame()
				games[t.GetRoomId()] = g
				t.PutResponse("Create Success")
			default:
				t.PutResponse(errors.New("Unknown event"))
			}
		}
	}
}
