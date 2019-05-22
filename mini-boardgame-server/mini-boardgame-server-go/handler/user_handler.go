package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUser(c echo.Context) error {
	if user, err := GetCurrentUser(c); err == nil {
		return c.JSON(http.StatusOK, J{
			"username": user.Username,
			"role":     user.GetRoleStrings(),
		})
	} else {
		return err
	}
}
