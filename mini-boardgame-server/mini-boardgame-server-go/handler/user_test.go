package handler

import (
	"net/http"
	"strconv"
	"testing"
)

func TestGetUser(t *testing.T) {
	TestHttp{
		test:    t,
		handler: GetUser,
		response: Response{
			Body: J{
				"id":       USER.ID,
				"username": USER.Username,
				"role":     USER.GetRoleStrings(),
			},
		},
		setups: []Setup{
			WithLogin(t, USER),
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
			Body: J{
				"id":       USER.ID,
				"username": USER.Username,
				"role":     USER.GetRoleStrings(),
			},
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(t, USER),
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
			Body: J{
				"id":       ADMIN.ID,
				"username": ADMIN.Username,
				"role":     ADMIN.GetRoleStrings(),
			},
		},
		setups: []Setup{
			WithUser(t, ADMIN),
			WithLogin(t, USER),
		},
	}.Run()
}
