package model

import (
	"time"
)

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
	return r.FindPlayer(func(p *Player) bool {
		return p.Seat == seat
	})
}

func (r *Room) FindPlayer(f func(p *Player) bool) (*Player, bool) {
	for _, p := range r.Players {
		if f(p) {
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
	if len(r.Players) != int(r.PlayerCount) {
		return false
	}
	for _, p := range r.Players {
		if p.State == NOT_READY {
			return false
		}
	}
	return true
}

func (r *Room) FindSeatByPlayer(userId uint) int {
	for _, v := range r.Players {
		if v.UserID == userId {
			return int(v.Seat)
		}
	}
	return -1
}

func (r *Room) GetNextSeat() (uint, bool) {
	for i := uint(0); i < r.PlayerCount; i++ {
		if _, ok := r.FindPlayerBySeat(i); !ok {
			return i, true
		}
	}
	return 0, false
}

func (r *Room) RemovePlayer(userId uint) {
	new := make([]*Player, 0)
	for _, v := range r.Players {
		if v.UserID != userId {
			new = append(new, v)
		} else {
			v.Room = nil
			v.RoomID = 0
			v.State = OUT_OF_GAME
		}
	}
	r.Players = new
}
