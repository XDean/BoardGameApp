package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRoom(t *testing.T) {
	TestHttp{
		test:    t,
		handler: CreateRoom,
		request: Request{
			Body: xecho.J{
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
				assert.Equal(t, "game name", room.GameId)
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
			Body: xecho.J{
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
			Body: xecho.J{
				"ID":          ROOMID,
				"GameId":      ROOM.GameId,
				"RoomName":    ROOM.RoomName,
				"PlayerCount": ROOM.PlayerCount,
				"Players": []xecho.J{
					{
						"UserID":      USERID,
						"State":       model.HOST,
						"StateString": model.HOST.String(),
						"Seat":        0,
					},
				},
			},
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
