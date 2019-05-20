package model

import (
	"time"
)

type Room struct {
	ID          uint `gorm:"primary_key"`
	GameName    string
	PlayerCount uint8
	RoomName    string
	CreatedTime time.Time
	Players     []Player
}

type Player struct {
	ID     uint `gorm:"primary_key"`
	UserID uint
	RoomID uint
	Seat   uint8
	Ready  bool
}
