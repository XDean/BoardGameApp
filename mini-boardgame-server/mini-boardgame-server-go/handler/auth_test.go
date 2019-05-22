package handler

import (
	"github.com/XDean/MiniBoardgame/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSignUp(t *testing.T) {

	testHttp(t, SignUp, Request{
		Body: J{
			"username": "username",
			"password": "password123",
		},
	}, Response{
		Code: http.StatusCreated,
	})

	testHttp(t, SignUp, Request{
		Body: J{
			"something": "wrong",
		},
	}, Response{
		Code:  http.StatusBadRequest,
		Error: true,
	})

	testHttp(t, SignUp, Request{
		Body: J{
			"username": "_",
			"password": "@#$",
		},
	}, Response{
		Code:  http.StatusBadRequest,
		Error: true,
	})
}

func TestLogin(t *testing.T) {
	user := &model.User{
		Username: "username",
		Password: "pwd123456",
	}
	err := user.CreateAccount()
	assert.NoError(t, err)

	testHttp(t, Login, Request{
		Body: J{
			"username": "username",
			"password": "pwd123456",
		},
	}, Response{})

	testHttp(t, Login, Request{
		Body: J{
			"username": "wrong",
			"password": "pwd123456",
		},
	}, Response{
		Code:  http.StatusUnauthorized,
		Error: true,
	})

	testHttp(t, Login, Request{
		Body: J{
			"username": "username",
			"password": "wrong",
		},
	}, Response{
		Code:  http.StatusUnauthorized,
		Error: true,
	})

	testHttp(t, Login, Request{
		Body: J{
			"wrong": "wrong",
		},
	}, Response{
		Code:  http.StatusBadRequest,
		Error: true,
	})
}
