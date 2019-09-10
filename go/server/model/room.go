package model

import "time"

type Room struct {
	ID          uint `gorm:"primary_key"`
	GameId      string
	PlayerCount uint
	RoomName    string
	CreatedTime time.Time
	Players     []*Player
	Options     *StringMap `gorm:"type:json"`
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

func (r *Room) FindHost() *Player {
	for _, p := range r.Players {
		if p.State == HOST {
			return p
		}
	}
	return nil
}

func (r *Room) IsAllReady() bool {
	for _, p := range r.Players {
		if p.State == NOT_READY {
			return false
		}
	}
	return true
}
