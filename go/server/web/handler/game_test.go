package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/games/rps"
	"github.com/xdean/miniboardgame/go/server/game"
	"github.com/xdean/miniboardgame/go/server/web/middleware"
	"net/http"
	"testing"
)

func init() {
	game.Register(rps.Instance)
}

func TestGetGameList(t *testing.T) {
	TestHttp{
		test:    t,
		handler: GetGameList,
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
					Body: xecho.M("Create Success"),
				},
				middleware: []echo.MiddlewareFunc{middleware.AuthRoom()},
				setups: []Setup{
					WithLogin(t, USER),
				},
			},
		},
	}.Run()
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
				handler: StartGame,
				response: Response{
					Error: true,
					Code:  http.StatusBadRequest,
					Body:  xecho.M("Players not ready"),
				},
				middleware: []echo.MiddlewareFunc{middleware.AuthRoom()},
				setups: []Setup{
					WithLogin(t, USER),
				},
			},
		},
	}.Run()
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
					Error: true,
					Code:  http.StatusBadRequest,
					Body:  xecho.M("You are not host"),
				},
				middleware: []echo.MiddlewareFunc{middleware.AuthRoom()},
				setups: []Setup{
					WithLogin(t, USER2),
				},
			},
		},
	}.Run()
}
