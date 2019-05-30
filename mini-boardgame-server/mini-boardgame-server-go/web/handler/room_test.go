package handler

import (
	"github.com/XDean/MiniBoardgame/model"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRoom(t *testing.T) {
	TestHttp{
		test:    t,
		handler: CreateRoom,
		request: Request{
			Body: J{
				"game_name":    "game name",
				"room_name":    "room name",
				"player_count": 3,
			},
		},
		response: Response{
			Extra: func(db *gorm.DB, recorder *httptest.ResponseRecorder) {
				room := new(model.Room)
				err := room.FindByUserID(db, USERID)
				assert.NoError(t, err)
				assert.Equal(t, "game name", room.GameName)
				assert.Equal(t, "room name", room.RoomName)
				assert.Equal(t, uint(3), room.PlayerCount)
				assert.Equal(t, uint(USERID), room.Players[0].UserID)
			},
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(t, USER),
		},
	}.Run()
}

func TestCreateRoomExist(t *testing.T) {
	TestHttp{
		test:    t,
		handler: CreateRoom,
		request: Request{
			Body: J{
				"game_name":    "game name",
				"room_name":    "room name",
				"player_count": 3,
			},
		},
		response: Response{
			Code:  http.StatusMethodNotAllowed,
			Error: true,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(t, USER),
			WithRoom(t, ROOM),
		},
	}.Run()
}

func TestGetRoom(t *testing.T) {
	TestHttp{
		test:    t,
		handler: GetRoom,
		response: Response{
			Body: ROOM,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(t, USER),
			WithRoom(t, ROOM),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetRoom,
		response: Response{
			Code:  http.StatusNotFound,
			Error: true,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(t, USER),
		},
	}.Run()
}
