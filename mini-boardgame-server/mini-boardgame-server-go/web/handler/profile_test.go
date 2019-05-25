package handler

import (
	"github.com/XDean/MiniBoardgame/model"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
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
			Code:  http.StatusUnauthorized,
			Error: true,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithProfile(t, USER_PROFILE),
		},
	}.Run()
}

func TestGetProfileById(t *testing.T) {
	TestHttp{
		test:    t,
		handler: GetProfileById,
		request: Request{
			Params: Params{
				"id": strconv.Itoa(USERID),
			},
		},
		response: Response{
			Body: USER_PROFILE,
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(t, USER),
			WithProfile(t, USER_PROFILE),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetProfileById,
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
			WithUser(t, USER),
			WithLogin(t, USER),
			WithProfile(t, USER_PROFILE),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetProfileById,
		request: Request{
			Params: Params{
				"id": strconv.Itoa(USERID),
			},
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(t, USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: GetProfileById,
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
			WithUser(t, USER),
			WithLogin(t, USER),
		},
	}.Run()
}

func TestUpdateProfile(t *testing.T) {
	TestHttp{
		test:    t,
		handler: UpdateProfile,
		request: Request{
			Body: J{
				"nickname":  "new",
				"sex":       model.Female,
				"avatarurl": "new-url",
			},
		},
		response: Response{
			Code: http.StatusAccepted,
			Extra: func(db *gorm.DB, recorder *httptest.ResponseRecorder) {
				profile := new(model.Profile)
				err := profile.FindByUserID(db, USERID)
				assert.NoError(t, err)
				assert.Equal(t, "new", profile.Nickname)
				assert.Equal(t, model.Female, profile.Sex)
				assert.Equal(t, "new-url", profile.AvatarURL)
			},
		},
		setups: []Setup{
			WithUser(t, USER),
			WithLogin(t, USER),
			WithProfile(t, USER_PROFILE),
		},
	}.Run()
}
