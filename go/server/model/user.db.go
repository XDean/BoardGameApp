package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (user *User) FindByID(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Find(user).Error
}

func (user *User) FindByUsername(db *gorm.DB, username string) error {
	return db.Where("username = ?", username).Find(user).Error
}

func (user *User) save(db *gorm.DB) error {
	return db.Save(user).Error
}

func (user *User) CreateAccount(db *gorm.DB) error {
	if encoded, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10); err == nil {
		user.Password = string(encoded)
		result := db.FirstOrCreate(user, User{Username: user.Username})
		if result.Error != nil {
			return result.Error
		} else if result.RowsAffected > 0 {
			return nil
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Username existed")
		}
	} else {
		return err
	}
}

func (user *User) ChangePassword(db *gorm.DB, old, new string) error {
	if !user.MatchPassword(old) {
		return errors.New("Password not correct")
	}
	if newPassword, err := bcrypt.GenerateFromPassword([]byte(new), 10); err == nil {
		user.Password = string(newPassword)
		return user.save(db)
	} else {
		return err
	}
}

func UserExistById(db *gorm.DB, id uint) (bool, error) {
	user := new(User)
	err := user.FindByID(db, id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func UserExistByUsername(db *gorm.DB, username string) (bool, error) {
	user := new(User)
	err := user.FindByUsername(db, username)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}
