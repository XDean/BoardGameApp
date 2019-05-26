package model

import "time"

type Room struct {
	ID          uint `gorm:"primary_key"`
	GameName    string
	PlayerCount uint
	RoomName    string
	CreatedTime time.Time
	Players     []*Player
}

func (r *Room) normalize() {
	for _, player := range r.Players {
		player.Room = r
		player.normalize()
	}
}
