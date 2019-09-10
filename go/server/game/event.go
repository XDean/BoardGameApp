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

	CreateRoomEvent struct {
		BaseEvent
		Options map[string]string
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
