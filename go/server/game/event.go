package game

type (
	Event interface {
		PutResponse(Response)
		GetResponse() Response
		GetPlayerId() int
		GetRoomId() int
	}

	BaseEvent struct {
		ResponseStream chan Response
		PlayerId       int
		RoomId         int
	}

	Response interface {
	}

	NewGameEvent struct {
		BaseEvent
		Options     map[string]string
		EventStream chan<- interface{}
	}
)

func (e BaseEvent) PutResponse(res Response) {
	e.ResponseStream <- res
}

func (e BaseEvent) GetResponse() Response {
	return <-e.ResponseStream
}

func (e BaseEvent) GetPlayerId() int {
	return e.PlayerId
}

func (e BaseEvent) GetRoomId() int {
	return e.RoomId
}
