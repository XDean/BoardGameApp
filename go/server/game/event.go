package game

import "github.com/xdean/miniboardgame/go/server/model"

type (
	Event interface {
		PutResponse(Response)
		GetResponse() Response
		SetUser(user *model.User)
		GetUser() *model.User
		SetRoom(room *model.Room)
		GetRoom() *model.Room
	}

	BaseEvent struct {
		ResponseStream chan Response
		User           *model.User
		Room           *model.Room
	}

	Response interface {
	}
)

func (e *BaseEvent) SetUser(user *model.User) {
	e.User = user
}

func (e *BaseEvent) SetRoom(room *model.Room) {
	e.Room = room
}

func (e *BaseEvent) PutResponse(res Response) {
	e.ResponseStream <- res
	close(e.ResponseStream)
}

func (e *BaseEvent) GetResponse() Response {
	return <-e.ResponseStream
}

func (e *BaseEvent) GetUser() *model.User {
	return e.User
}

func (e *BaseEvent) GetRoom() *model.Room {
	return e.Room
}

func (e *BaseEvent) GetSeat() int {
	return e.Room.FindSeatByPlayer(e.User.ID)
}
