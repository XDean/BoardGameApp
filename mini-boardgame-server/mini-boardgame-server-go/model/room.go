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

func (r *Room) FindPlayerBySeat(seat uint) (*Player, bool) {
	for _, p := range r.Players {
		if seat == p.Seat {
			return p, true
		}
	}
	return nil, false
}
