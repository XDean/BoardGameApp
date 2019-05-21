package handler

import (
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SignUp(c echo.Context) error {
	type Param struct {
		Username string `json:"username" validate:"required,regexp=USERNAME"`
		Password string `json:"password" validate:"required,regexp=PASSWORD"`
	}
	param := new(Param)
	if err := c.Bind(param); err != nil {
		return err
	}
	if err := c.Validate(param); err != nil {
		return err
	}
	user := &model.User{
		Username: param.Username,
		Password: param.Password,
		Roles:    []model.Role{{Name: _const.ROLE_USER}},
	}
	return user.CreateAccount()
}

func Login(c echo.Context) error {
	type Param struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	param := new(Param)
	if err := c.Bind(param); err != nil {
		return err
	}
	if err := c.Validate(param); err != nil {
		return err
	}
	user := new(model.User)
	if err := user.FindByUsername(param.Username); err == nil {
		if user.MatchPassword(param.Password) {
			return c.JSON(http.StatusOK, _const.H{})
		}
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "Bad Credentials")
}
