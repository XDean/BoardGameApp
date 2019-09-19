package game

import "github.com/xdean/miniboardgame/go/server/model"

type (
	Event interface {
		PutResponse(Response)
		GetResponse() Response
		GetUser() *model.User
		GetRoom() *model.Room
	}

	BaseEvent struct {
		ResponseStream chan Response
		User           *model.User
		Room           *model.Room
	}

	Response interface {
	}

	NewGameEvent struct {
		BaseEvent
	}
)

func (e BaseEvent) PutResponse(res Response) {
	e.ResponseStream <- res
	close(e.ResponseStream)
}

func (e BaseEvent) GetResponse() Response {
	return <-e.ResponseStream
}

func (e BaseEvent) GetUser() *model.User {
	return e.User
}

func (e BaseEvent) GetRoom() *model.Room {
	return e.Room
}
