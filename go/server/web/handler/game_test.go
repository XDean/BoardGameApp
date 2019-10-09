package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/model"
	"net/http/httptest"
	"testing"
)

func TestGetGameList(t *testing.T) {
	TestHttp{
		test:    t,
		handler: GetGameList,
		response: Response{
			Body: xecho.J{
				"Games": []xecho.J{
					{
						"Id":   GAME_ID,
						"Name": GAME_NAME,
						"Player": xecho.J{
							"Min": 2,
							"Max": 3,
						},
					},
				},
			},
		},
	}.Run()
}

func TestStartGame(t *testing.T) {
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
			WithLogin(t, USER),
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
			WithLogin(t, USER2),
		},
	}
	user2Ready := TestHttp{
		handler: Ready,
		setups: []Setup{
			WithLogin(t, USER2),
		},
	}
	TestHttpSeries{
		test: t,
		setups: []Setup{
			WithUser(t, USER),
			WithUser(t, USER2),
		},
		children: []TestHttp{
			user1Create,
			user2Join,
			user2Ready,
			{
				handler: StartGame,
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
					WithLogin(t, USER),
				},
			},
		},
	}.Run()
}
