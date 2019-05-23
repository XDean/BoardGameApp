package handler

import (
	"net/http"
	"testing"
)

func TestSignUp(t *testing.T) {
	TestHttp{
		test:    t,
		handler: SignUp,
		req: Request{
			Body: J{
				"username": USERNAME,
				"password": USERPWD,
			},
		},
		res: Response{
			Code: http.StatusCreated,
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: SignUp,
		req: Request{
			Body: J{
				"something": "wrong",
			},
		},
		res: Response{
			Code:  http.StatusBadRequest,
			Error: true,
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: SignUp,
		req: Request{
			Body: J{
				"username": "_",
				"password": "@#$",
			},
		},
		res: Response{
			Code:  http.StatusBadRequest,
			Error: true,
		},
	}.Run()
}

func TestLogin(t *testing.T) {
	TestHttp{
		test:    t,
		handler: Login,
		req: Request{
			Body: J{
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
		req: Request{
			Body: J{
				"username": "wrong",
				"password": "pwd123456",
			},
		},
		res: Response{
			Code:  http.StatusUnauthorized,
			Error: true,
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: Login,
		req: Request{
			Body: J{
				"username": "username",
				"password": "wrong",
			},
		},
		res: Response{
			Code:  http.StatusUnauthorized,
			Error: true,
		},
	}.Run()

	TestHttp{
		test:    t,
		handler: Login,
		req: Request{
			Body: J{
				"wrong": "wrong",
			},
		},
		res: Response{
			Code:  http.StatusBadRequest,
			Error: true,
		},
	}.Run()
}
