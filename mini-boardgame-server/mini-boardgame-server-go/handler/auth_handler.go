package handler

import (
	"github.com/XDean/MiniBoardgame/config"
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
	if err := user.CreateAccount(GetDB(c)); err == nil {
		if t, err := user.GenerateToken(config.Global.Security.Key); err == nil {
			return c.JSON(http.StatusCreated, J{
				"message": "Sign up success",
				"token":   t,
			})
		} else {
			return err
		}
	} else {
		return err
	}
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
	if err := user.FindByUsername(GetDB(c), param.Username); err == nil {
		if user.MatchPassword(param.Password) {
			if t, err := user.GenerateToken(config.Global.Security.Key); err == nil {
				return c.JSON(http.StatusOK, J{
					"message": "Login success",
					"token":   t,
				})
			} else {
				return err
			}
		}
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "Bad Credentials")
}
