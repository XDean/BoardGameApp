package model

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique;not null"`
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
	UserID    uint `gorm:"unique;not null"`
	User      User
	Nickname  string
	Male      bool
	AvatarURL string
}

func (user *User) GetRoleStrings() []string {
	roles := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = role.Name
	}
	return roles
}

func (user *User) SetRoles(roles []string) {
	user.Roles = make([]Role, len(roles))
	for i, role := range roles {
		user.Roles[i] = Role{
			UserID: user.ID,
			Name:   role,
		}
	}
}

func (user *User) MatchPassword(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	return err == nil
}

func (user *User) GenerateToken(key string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID:   user.ID,
		Username: user.Username,
		Roles:    user.GetRoleStrings(),
		StandardClaims: jwt.StandardClaims{
			Subject:   "access token",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}).SignedString([]byte(key))
}
