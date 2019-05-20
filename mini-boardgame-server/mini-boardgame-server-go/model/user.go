package model

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
