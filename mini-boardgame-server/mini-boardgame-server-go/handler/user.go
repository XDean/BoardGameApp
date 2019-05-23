package handler

import (
	"github.com/XDean/MiniBoardgame/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

func GetUserById(c echo.Context) error {
	idParam := c.Param("id")
	if id, err := strconv.Atoi(idParam); err == nil {
		user := new(model.User)
		if err := user.FindByID(GetDB(c), uint(id)); err == nil {
			return c.JSON(http.StatusOK, J{
				"username": user.Username,
				"role":     user.GetRoleStrings(),
			})
		} else {
			return err
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Unrecognized ID: "+idParam)
	}
}
