package handler

import (
	"net/http"
	"testing"
)

func TestSignUp(test *testing.T) {
	testHttp(test, Request{
		Body: J{
			"username": "username",
			"password": "password123",
		},
	}, Response{
		Code: http.StatusCreated,
	}, SignUp)

	testHttp(test, Request{
		Body: J{
			"something": "wrong",
		},
	}, Response{
		Code:  http.StatusBadRequest,
		Error: true,
	}, SignUp)

	testHttp(test, Request{
		Body: J{
			"username": "_",
			"password": "@#$",
		},
	}, Response{
		Code:  http.StatusBadRequest,
		Error: true,
	}, SignUp)
}
