package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Room struct {
	gorm.Model
	GameName    string
	PlayerCount uint8
	RoomName    string
	CreatedTime time.Time
	Players     []Player
}

type Player struct {
	gorm.Model
	UserID uint
	RoomID uint
	Seat   uint8
	Ready  bool
}
