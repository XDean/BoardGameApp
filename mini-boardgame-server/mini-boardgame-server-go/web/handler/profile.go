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
		profile := model.EmptyProfile(user.ID)
		if err := profile.FindByUserID(GetDB(c), user.ID); err == nil || gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusOK, profile)
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
			profile := model.EmptyProfile(user.ID)
			if err := profile.FindByUserID(GetDB(c), user.ID); err == nil || gorm.IsRecordNotFoundError(err) {
				return c.JSON(http.StatusOK, profile)
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

func UpdateProfile(c echo.Context) error {
	type Param struct {
		Nickname  string    `json:"nickname" query:"nickname" form:"nickname"`
		Sex       model.Sex `json:"sex" query:"sex" form:"sex"`
		AvatarURL string    `json:"avatarurl" query:"avatarurl" form:"avatarurl"`
	}
	param := new(Param)
	if err := c.Bind(param); err != nil {
		return err
	}
	if err := c.Validate(param); err != nil {
		return err
	}
	if user, err := GetCurrentUser(c); err == nil {
		profile := model.EmptyProfile(user.ID)
		if err := profile.FindByUserID(GetDB(c), user.ID); err == nil || gorm.IsRecordNotFoundError(err) {
			if param.Nickname != "" {
				profile.Nickname = param.Nickname
			}
			if param.AvatarURL != "" {
				profile.AvatarURL = param.AvatarURL
			}
			if param.Sex != model.Unknown {
				profile.Sex = param.Sex
			}
			if err := profile.Save(GetDB(c)); err == nil {
				return c.JSON(http.StatusOK, profile)
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}
