package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/model"
	"net/http"
	"strconv"
)

func GetProfile(c echo.Context) error {
	user, err := GetCurrentUser(c)
	xecho.MustNoError(err)
	profile := model.EmptyProfile(user.ID)
	err = profile.FindByUserID(GetDB(c), user.ID)
	if err == nil || gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusOK, profile)
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
	xecho.MustBindAndValidate(c, param)
	user, err := GetCurrentUser(c)
	xecho.MustNoError(err)
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
		err := profile.Save(GetDB(c))
		xecho.MustNoError(err)
		return c.JSON(http.StatusOK, profile)
	} else {
		return err
	}
}
