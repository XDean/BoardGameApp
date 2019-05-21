package model

import (
	"errors"
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint `gorm:"primary_key"`
	Username string
	Password string
	Roles    []Role
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
		return user, user.FindByID(userID)
	}
	return nil, errors.New("not authorized")
}

func (user *User) FindByID(id uint) error {
	if err := DB.Where("id = ?", id).Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) FindByUsername(username string) error {
	if err := DB.Where("username = ?", username).Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) Save() error {
	return DB.Save(user).Error
}

func (profile *Profile) Save() error {
	return DB.Save(profile).Error
}

func (user *User) CreateAccount() error {
	if encoded, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10); err == nil {
		user.Password = string(encoded)
		return user.Save()
	} else {
		return err
	}
}

func (user *User) MatchPassword(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	return err == nil
}

func (user *User) ChangePassword(old, new string) error {
	if user.MatchPassword(old) {
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
	if err := user.FindByID(id); gorm.IsRecordNotFoundError(err) {
		return false, nil
	} else {
		return false, err
	}
	return true, nil
}

func UserExistByUsername(username string) (bool, error) {
	user := new(User)
	if err := user.FindByUsername(username); gorm.IsRecordNotFoundError(err) {
		return false, nil
	} else {
		return false, err
	}
	return true, nil
}
