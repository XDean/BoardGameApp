package handler

import (
	"github.com/XDean/MiniBoardgame/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetProfile(c echo.Context) error {
	if user, err := GetCurrentUser(c); err == nil {
		profile := new(model.Profile)
		if err := profile.FindByUserID(GetDB(c), user.ID); err != nil {
			return c.JSON(http.StatusOK, *profile)
		} else if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNoContent, model.EmptyProfile(user.ID))
		} else {
			return err
		}
	} else {
		return err
	}
}

func GetProfileById(c echo.Context) error {
	idParam := c.Param("id")
	if id, err := strconv.Atoi(idParam); err == nil {
		user := new(model.User)
		if err := user.FindByID(GetDB(c), uint(id)); err == nil {
			profile := new(model.Profile)
			if err := profile.FindByUserID(GetDB(c), user.ID); err != nil {
				return c.JSON(http.StatusOK, *profile)
			} else if gorm.IsRecordNotFoundError(err) {
				return c.JSON(http.StatusNoContent, model.EmptyProfile(user.ID))
			} else {
				return err
			}
		} else {
			return DBNotFound(err, "No such user")
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Unrecognized id: "+idParam)
	}
}
