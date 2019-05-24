package handler

import "net/http"

type ErrorCode struct {
	Status int
	Code   string
}

var (
	BAD_CREDENTIALS = Error(http.StatusUnauthorized, "BAD_CREDENTIALS")
)

func Error(status int, code string) ErrorCode {
	return ErrorCode{
		Status: status,
		Code:   code,
	}
}
