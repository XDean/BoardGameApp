package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Enabled  bool
	Roles    []Role
	Profile  Profile
}

type Role struct {
	gorm.Model
	Name   string
	UserID uint
}

type Profile struct {
	gorm.Model
	Nickname  string
	Male      bool
	AvatarURL string
}
