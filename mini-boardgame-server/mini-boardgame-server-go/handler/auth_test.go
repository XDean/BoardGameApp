package handler

import (
	"github.com/XDean/MiniBoardgame/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"testing"
)

func TestSignUp(t *testing.T) {
	TestHttp{
		test:    t,
		handler: SignUp,
		request: Request{
			Body: J{
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
			Body: J{
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
			Body: J{
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
			Body: J{
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
			Body: J{
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
			Body: J{
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
			Body: J{
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
			Body: J{
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
	TestHttp{
		test:    t,
		handler: Login,
		request: Request{
			Body: J{
				"type":     "openid",
				"provider": "test",
				"token":    "token",
			},
		},
		setups: []Setup{
			WithOpenid(),
		},
	}.Run()
	TestHttp{
		test:    t,
		handler: Login,
		request: Request{
			Body: J{
				"type":     "openid",
				"provider": "test",
				"token":    "token",
			},
		},
		setups: []Setup{
			WithOpenid(),
			WithUser(t, &model.User{
				Username: "token@test",
			}),
		},
	}.Run()
}
