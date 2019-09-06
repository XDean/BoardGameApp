package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/model"
	"net/http"
	"strconv"
)

func GetUser(c echo.Context) error {
	if user, err := GetCurrentUser(c); err == nil {
		return c.JSON(http.StatusOK, xecho.J{
			"id":       user.ID,
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
			return c.JSON(http.StatusOK, xecho.J{
				"id":       user.ID,
				"username": user.Username,
				"role":     user.GetRoleStrings(),
			})
		} else {
			return DBNotFound(err, "No such user")
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Unrecognized id: "+idParam)
	}
}
