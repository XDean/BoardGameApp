package handler

import (
	"github.com/XDean/MiniBoardgame/model"
	"net/http"
	"testing"
)

func TestGetProfile(t *testing.T) {
	TestHttp{
		test:    t,
		handler: GetProfile,
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(t, USER),
			WithProfile(t, USER_PROFILE),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetProfile,
		response: Response{
			Code: http.StatusNoContent,
			Body: model.EmptyProfile(USERID),
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(t, USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetProfile,
		response: Response{
			Code: http.StatusUnauthorized,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithProfile(t, USER_PROFILE),
		},
	}.Run()
}
