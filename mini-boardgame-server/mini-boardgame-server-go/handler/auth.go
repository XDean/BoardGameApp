package handler

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/validator.v2"
)

type SignUpParam struct {
	Username string `validate:"min=6,regexp=^(?!_)(?!.*?_$)[a-zA-Z0-9_]+$"`
	Password string `validate:"min=6,regexp=^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$"`
}

func SignUp(c echo.Context) error {
	param := new(SignUpParam)
	if err := c.Bind(param); err != nil {
		return err
	}
	if err := validator.Validate(param); err != nil {
		return err
	}
}
