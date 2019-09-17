package model

import (
	"fmt"
	"github.com/xdean/miniboardgame/go/server/model/space"
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

func (r *Room) EventHostId() string {
	return fmt.Sprintf("ROOM-%d", r.ID)
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

func (r *Room) SendEvent(e space.Event) {
	space.SendEvent(r, e)
}

func (r *Room) Listen() space.Subscription {
	return space.Listen(r)
}
