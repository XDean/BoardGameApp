package model

import (
	"errors"
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/db"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint `gorm:"primary_key"`
	Username string
	Password string
	Roles    []Role
	Profile  Profile
}

type Role struct {
	ID     uint `gorm:"primary_key"`
	UserID uint
	Name   string
}

type Profile struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	User      User
	Nickname  string
	Male      bool
	AvatarURL string
}

func GetCurrentUser(c echo.Context) (*User, error) {
	if user, ok := c.Get(_const.USER_ENTITY).(*User); ok {
		return user, nil
	}
	if userID, ok := c.Get(_const.USERID).(uint); ok {
		user := new(User)
		return user, user.GetByID(userID)
	}
	return nil, errors.New("not authorized")
}

func (user *User) GetByID(id uint) error {
	if err := db.DB.Where("id = ?", id).Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) GetByUsername(username string) error {
	if err := db.DB.Where("username = ?", username).Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) Save() error {
	return db.DB.Save(user).Error
}

func (profile *Profile) Save() error {
	return db.DB.Save(profile).Error
}

func (user *User) ChangePassword(old, new string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(old)); err != nil {
		return errors.New("Password not correct")
	}
	if newPassword, err := bcrypt.GenerateFromPassword([]byte(new), 10); err == nil {
		user.Password = string(newPassword)
		return user.Save()
	} else {
		return err
	}
}

func UserExistById(id uint) (bool, error) {
	user := new(User)
	if err := user.GetByID(id); gorm.IsRecordNotFoundError(err) {
		return false, nil
	} else {
		return false, err
	}
	return true, nil
}

func UserExistByUsername(username string) (bool, error) {
	user := new(User)
	if err := user.GetByUsername(username); gorm.IsRecordNotFoundError(err) {
		return false, nil
	} else {
		return false, err
	}
	return true, nil
}
