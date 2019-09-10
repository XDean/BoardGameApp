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
			event.PutResponse(handleEvent(event))
		}
	}
}

func handleEvent(event game.Event) game.Response {
	room := games[event.GetRoomId()]
	switch t := event.(type) {
	case Event:
		if room == nil {
			return errors.New("No such game room")
		}
		g := games[event.GetRoomId()]
		return g.Play(t)
	case game.NewGameEvent:
		if room != nil {
			return errors.New("The game has started")
		}
		g := NewStandardGame()
		g.EventStream = t.EventStream
		games[event.GetRoomId()] = g
		return "Create Success"
	default:
		return errors.New("Unknown event")
	}
}
