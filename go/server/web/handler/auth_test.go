package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/model"
	"github.com/xdean/miniboardgame/go/server/web/handler/openid"
	"net/http"
	"testing"
)

func TestSignUp(t *testing.T) {
	TestHttp{
		test:    t,
		handler: SignUp,
		request: Request{
			Body: xecho.J{
				"username": USERNAME,
				"password": USERPWD,
			},
		},
		response: Response{
			Code: http.StatusCreated,
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: SignUp,
		request: Request{
			Body: xecho.J{
				"something": "wrong",
			},
		},
		response: Response{
			Code:  http.StatusBadRequest,
			Error: true,
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: SignUp,
		request: Request{
			Body: xecho.J{
				"username": "_",
				"password": "@#$",
			},
		},
		response: Response{
			Code:  http.StatusBadRequest,
			Error: true,
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: SignUp,
		request: Request{
			Method: echo.POST,
		},
		response: Response{
			Code:  http.StatusBadRequest,
			Error: true,
		},
	}.Run()
}

func TestLogin(t *testing.T) {
	TestHttp{
		test:    t,
		handler: Login,
		request: Request{
			Body: xecho.J{
				"type": "wrong",
			},
		},
		response: Response{
			Code:  http.StatusBadRequest,
			Error: true,
		},
	}.Run()
}

func TestLoginPassword(t *testing.T) {
	TestHttp{
		test:    t,
		handler: Login,
		request: Request{
			Body: xecho.J{
				"type":     "password",
				"username": USERNAME,
				"password": USERPWD,
			},
		},
		setups: []Setup{
			WithUser(t, USER),
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: Login,
		request: Request{
			Body: xecho.J{
				"username": "wrong",
				"password": "pwd123456",
			},
		},
		response: Response{
			Code:  http.StatusUnauthorized,
			Error: true,
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: Login,
		request: Request{
			Body: xecho.J{
				"username": "username",
				"password": "wrong",
			},
		},
		response: Response{
			Code:  http.StatusUnauthorized,
			Error: true,
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: Login,
		request: Request{
			Body: xecho.J{
				"wrong": "wrong",
			},
		},
		response: Response{
			Code:  http.StatusBadRequest,
			Error: true,
		},
	}.Run()
}

func TestLoginOpenid(t *testing.T) {
	openid.Providers = map[string]openid.OpenIdProvider{
		"test": {
			Name: "test",
			Auth: func(token string) (string, error) {
				return token, nil
			},
		},
		"test-fail": {
			Name: "test",
			Auth: func(token string) (string, error) {
				return "", errors.New("openid fail")
			},
		},
	}

	TestHttp{
		test:    t,
		handler: Login,
		request: Request{
			Body: xecho.J{
				"type":     "openid",
				"provider": "test",
				"token":    "token",
			},
		},
	}.Run()
	TestHttp{
		test:    t,
		handler: Login,
		request: Request{
			Body: xecho.J{
				"type":     "openid",
				"provider": "test",
				"token":    "token",
			},
		},
		setups: []Setup{
			WithUser(t, &model.User{
				Username: "token@test",
			}),
		},
	}.Run()
}
