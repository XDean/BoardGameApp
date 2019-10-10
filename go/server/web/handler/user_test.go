package handler

import (
	"github.com/xdean/goex/xecho"
	"net/http"
	"strconv"
	"testing"
)

func TestGetUser(t *testing.T) {
	TestHttp{
		test:    t,
		handler: GetUser,
		response: Response{
			Body: xecho.J{
				"id":       USER.ID,
				"username": USER.Username,
				"role":     USER.GetRoleStrings(),
			},
		},
		setups: []Setup{
			WithLogin(USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetUser,
		response: Response{
			Code:  http.StatusUnauthorized,
			Error: true,
		},
		setups: []Setup{
			WithUser(t, USER),
		},
	}.Run()
}

func TestGetUserById(t *testing.T) {
	TestHttp{
		test:    t,
		handler: GetUserById,
		request: Request{
			Params: Params{
				"id": strconv.Itoa(USERID),
			},
		},
		response: Response{
			Body: xecho.J{
				"id":       USER.ID,
				"username": USER.Username,
				"role":     USER.GetRoleStrings(),
			},
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetUserById,
		request: Request{
			Params: Params{
				"id": strconv.Itoa(ADMINID),
			},
		},
		response: Response{
			Body: xecho.J{
				"id":       ADMIN.ID,
				"username": ADMIN.Username,
				"role":     ADMIN.GetRoleStrings(),
			},
		},
		setups: []Setup{
			WithUser(t, ADMIN),
			WithLogin(USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetUserById,
		request: Request{
			Params: Params{
				"id": "wrong",
			},
		},
		response: Response{
			Code:  http.StatusBadRequest,
			Error: true,
		},
		setups: []Setup{
			WithLogin(USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetUserById,
		request: Request{
			Params: Params{
				"id": strconv.Itoa(ADMINID),
			},
		},
		response: Response{
			Code:  http.StatusNotFound,
			Error: true,
		},
		setups: []Setup{
			WithLogin(USER),
		},
	}.Run()
}
