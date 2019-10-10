package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/model"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreateRoom(t *testing.T) {
	TestHttp{
		test:    t,
		handler: CreateRoom,
		request: Request{
			Body: xecho.J{
				"game_id":      GAME_ID,
				"room_name":    "room name",
				"player_count": 3,
			},
		},
		response: Response{
			Extra: func(db *gorm.DB, recorder *httptest.ResponseRecorder) {
				room := new(model.Room)
				err := room.FindByUserID(db, USERID)
				assert.NoError(t, err)
				assert.Equal(t, GAME_ID, room.GameId)
				assert.Equal(t, "room name", room.RoomName)
				assert.Equal(t, uint(3), room.PlayerCount)
				assert.Equal(t, uint(USERID), room.Players[0].UserID)
			},
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: CreateRoom,
		request: Request{
			Body: xecho.J{
				"game_id":      GAME_ID,
				"room_name":    "room name",
				"player_count": 100,
			},
		},
		response: Response{
			Error: true,
			Code:  http.StatusBadRequest,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: CreateRoom,
		request: Request{
			Body: xecho.J{
				"game_id":      "no game",
				"room_name":    "room name",
				"player_count": 5,
			},
		},
		response: Response{
			Error: true,
			Code:  http.StatusBadRequest,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(USER),
		},
	}.Run()
}

func TestCreateRoomExist(t *testing.T) {
	TestHttp{
		test:    t,
		handler: CreateRoom,
		request: Request{
			Body: xecho.J{
				"game_id":      GAME_ID,
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
			WithLogin(USER),
			WithCreateRoom(t, ROOM, USER),
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
			WithLogin(USER),
			WithCreateRoom(t, ROOM, USER),
			WithInRoom(ROOM),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetRoomByID,
		request: Request{
			Params: Params{
				"id": strconv.Itoa(ROOMID),
			},
		},
		response: Response{
			Body: xecho.J{
				"ID":          ROOMID,
				"GameId":      ROOM.GameId,
				"RoomName":    ROOM.RoomName,
				"PlayerCount": ROOM.PlayerCount,
				"Players": []xecho.J{
					{
						"UserID":      USERID2,
						"State":       model.HOST,
						"StateString": model.HOST.String(),
						"Seat":        0,
					},
				},
			},
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(USER),
			WithCreateRoom(t, ROOM, USER2),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetRoomByID,
		request: Request{
			Params: Params{
				"id": "100",
			},
		},
		response: Response{
			Error: true,
			Code:  http.StatusNotFound,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetRoomByID,
		request: Request{
			Params: Params{
				"id": "a",
			},
		},
		response: Response{
			Error: true,
			Code:  http.StatusBadRequest,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetRoom,
		response: Response{
			Code:  http.StatusBadRequest,
			Error: true,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(USER),
		},
	}.Run()
}

func TestJoinRoom(t *testing.T) {
	user1Create := TestHttp{
		handler: CreateRoom,
		request: Request{
			Body: xecho.J{
				"game_id":      GAME_ID,
				"room_name":    "room name",
				"player_count": 2,
			},
		},
		response: Response{
			Extra: func(db *gorm.DB, recorder *httptest.ResponseRecorder) {
				room := new(model.Room)
				err := room.FindByUserID(db, USERID)
				assert.NoError(t, err)
				assert.Equal(t, GAME_ID, room.GameId)
				assert.Equal(t, "room name", room.RoomName)
				assert.Equal(t, uint(2), room.PlayerCount)
				assert.Equal(t, uint(USERID), room.Players[0].UserID)
			},
		},
		setups: []Setup{
			WithLogin(USER),
		},
	}
	user2Join := TestHttp{
		handler: JoinRoom,
		request: Request{
			Params: Params{
				"id": "1",
			},
		},
		response: Response{
			Extra: func(db *gorm.DB, recorder *httptest.ResponseRecorder) {
				player := new(model.Player)
				err := player.GetByUserID(db, USERID2)
				assert.NoError(t, err)
				assert.Equal(t, 2, len(player.Room.Players))
				assert.Equal(t, model.NOT_READY, player.State)
				assert.Equal(t, uint(1), player.RoomID)
				assert.Equal(t, uint(1), player.Seat)
			},
		},
		setups: []Setup{
			WithLogin(USER2),
		},
	}
	// 1 create 2 join 1 exit
	TestHttpSeries{
		test: t,
		setups: []Setup{
			WithUser(t, USER),
			WithUser(t, USER2),
		},
		children: []TestHttp{
			user1Create,
			user2Join,
			{
				handler: ExitRoom,
				response: Response{
					Extra: func(db *gorm.DB, recorder *httptest.ResponseRecorder) {
						player := new(model.Player)
						err := player.GetByUserID(db, USERID)
						assert.Nil(t, player.Room)
						assert.NoError(t, err)
						assert.Equal(t, model.OUT_OF_GAME, player.State)
						assert.Equal(t, uint(0), player.RoomID)
						assert.Equal(t, uint(0), player.Seat)

						player2 := new(model.Player)
						err = player2.GetByUserID(db, USERID2)
						assert.NoError(t, err)
						assert.NotNil(t, player2.Room)
						assert.Equal(t, 1, len(player2.Room.Players))
						assert.Equal(t, model.HOST, player2.State)
						assert.Equal(t, uint(1), player2.RoomID)
						assert.Equal(t, uint(1), player2.Seat)
					},
				},
				setups: []Setup{
					WithLogin(USER),
				},
			},
		},
	}.Run()
	// 1 create 1 join
	TestHttpSeries{
		test: t,
		setups: []Setup{
			WithUser(t, USER),
			WithUser(t, USER2),
		},
		children: []TestHttp{
			user1Create,
			{
				handler: JoinRoom,
				request: Request{
					Params: Params{
						"id": "1",
					},
				},
				response: Response{
					Code: http.StatusBadRequest,
				},
				setups: []Setup{
					WithLogin(USER),
				},
			},
		},
	}.Run()
	// 1 create 2 join 3 join
	TestHttpSeries{
		test: t,
		setups: []Setup{
			WithUser(t, USER),
			WithUser(t, USER2),
		},
		children: []TestHttp{
			user1Create,
			user2Join,
			{
				handler: JoinRoom,
				request: Request{
					Params: Params{
						"id": "1",
					},
				},
				response: Response{
					Code:  http.StatusBadRequest,
					Error: true,
				},
				setups: []Setup{
					WithLogin(USER3),
				},
			},
		},
	}.Run()
	// 1 create 2 exit
	TestHttpSeries{
		test: t,
		setups: []Setup{
			WithUser(t, USER),
			WithUser(t, USER2),
		},
		children: []TestHttp{
			user1Create,
			{
				handler: ExitRoom,
				response: Response{
					Code: http.StatusBadRequest,
				},
				setups: []Setup{
					WithLogin(USER2),
				},
			},
		},
	}.Run()
}

func TestSwapSeat(t *testing.T) {
	user1Create := TestHttp{
		handler: CreateRoom,
		request: Request{
			Body: xecho.J{
				"game_id":      GAME_ID,
				"room_name":    "room name",
				"player_count": 3,
			},
		},
		setups: []Setup{
			WithLogin(USER),
		},
	}
	user2Join := TestHttp{
		handler: JoinRoom,
		request: Request{
			Params: Params{
				"id": "1",
			},
		},
		setups: []Setup{
			WithLogin(USER2),
		},
	}
	// swap 1 -> 2
	TestHttpSeries{
		test: t,
		setups: []Setup{
			WithUser(t, USER),
			WithUser(t, USER2),
		},
		children: []TestHttp{
			user1Create,
			user2Join,
			{
				handler: SwapSeat,
				request: Request{
					Params: Params{
						"seat": "2",
					},
				},
				response: Response{
					Extra: func(db *gorm.DB, recorder *httptest.ResponseRecorder) {
						player := new(model.Player)
						err := player.GetByUserID(db, USERID2)
						assert.NoError(t, err)
						assert.Equal(t, model.NOT_READY, player.State)
						assert.Equal(t, uint(1), player.RoomID)
						assert.Equal(t, uint(2), player.Seat)
					},
				},
				setups: []Setup{
					WithLogin(USER2),
				},
			},
		},
	}.Run()
	// swap 1 -> 1
	TestHttpSeries{
		test: t,
		setups: []Setup{
			WithUser(t, USER),
			WithUser(t, USER2),
		},
		children: []TestHttp{
			user1Create,
			user2Join,
			{
				handler: SwapSeat,
				request: Request{
					Params: Params{
						"seat": "1",
					},
				},
				response: Response{
					Code:  http.StatusBadRequest,
					Error: true,
				},
				setups: []Setup{
					WithLogin(USER2),
				},
			},
		},
	}.Run()
	// swap 1 <-> 0
	TestHttpSeries{
		test: t,
		setups: []Setup{
			WithUser(t, USER),
			WithUser(t, USER2),
		},
		children: []TestHttp{
			user1Create,
			user2Join,
			{
				handler: SwapSeat,
				request: Request{
					Params: Params{
						"seat": "0",
					},
				},
				response: Response{
					Code: http.StatusOK,
				},
				setups: []Setup{
					WithLogin(USER2),
				},
			},
			{
				handler: SwapSeat,
				request: Request{
					Params: Params{
						"seat": "1",
					},
				},
				response: Response{
					Code: http.StatusOK,
					Extra: func(db *gorm.DB, recorder *httptest.ResponseRecorder) {
						player := new(model.Player)
						err := player.GetByUserID(db, USERID)
						assert.NoError(t, err)
						assert.Equal(t, model.HOST, player.State)
						assert.Equal(t, uint(1), player.RoomID)
						assert.Equal(t, uint(1), player.Seat)

						player2 := new(model.Player)
						err = player2.GetByUserID(db, USERID2)
						assert.NoError(t, err)
						assert.Equal(t, model.NOT_READY, player2.State)
						assert.Equal(t, uint(1), player2.RoomID)
						assert.Equal(t, uint(0), player2.Seat)
					},
				},
				setups: []Setup{
					WithLogin(USER),
				},
			},
		},
	}.Run()
}

func TestReady(t *testing.T) {
	user1Create := TestHttp{
		handler: CreateRoom,
		request: Request{
			Body: xecho.J{
				"game_id":      GAME_ID,
				"room_name":    "room name",
				"player_count": 2,
			},
		},
		setups: []Setup{
			WithLogin(USER),
		},
	}
	user2Join := TestHttp{
		handler: JoinRoom,
		request: Request{
			Params: Params{
				"id": "1",
			},
		},
		setups: []Setup{
			WithLogin(USER2),
		},
	}
	// host ready
	TestHttpSeries{
		test: t,
		setups: []Setup{
			WithUser(t, USER),
			WithUser(t, USER2),
		},
		children: []TestHttp{
			user1Create,
			user2Join,
			{
				handler: Ready,
				response: Response{
					Error: true,
					Code:  http.StatusBadRequest,
				},
				setups: []Setup{
					WithLogin(USER),
				},
			},
		},
	}.Run()
	// host ready
	TestHttpSeries{
		test: t,
		setups: []Setup{
			WithUser(t, USER),
			WithUser(t, USER2),
		},
		children: []TestHttp{
			user1Create,
			user2Join,
			{
				handler: Ready,
				response: Response{
					Body: xecho.J{
						"ready": true,
					},
					Extra: func(db *gorm.DB, recorder *httptest.ResponseRecorder) {
						player := new(model.Player)
						err := player.GetByUserID(db, USERID2)
						assert.NoError(t, err)
						assert.Equal(t, model.READY, player.State)
					},
				},
				setups: []Setup{
					WithLogin(USER2),
				},
			},
			{
				handler: Ready,
				response: Response{
					Body: xecho.J{
						"ready": false,
					},
					Extra: func(db *gorm.DB, recorder *httptest.ResponseRecorder) {
						player := new(model.Player)
						err := player.GetByUserID(db, USERID2)
						assert.NoError(t, err)
						assert.Equal(t, model.NOT_READY, player.State)
					},
				},
				setups: []Setup{
					WithLogin(USER2),
				},
			},
		},
	}.Run()
}
