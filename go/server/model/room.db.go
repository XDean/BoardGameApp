package model

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xgo"
	"net/http"
	"time"
)

func (r *Room) save(db *gorm.DB) error {
	return db.Save(r).Error
}

func (r *Room) delete(db *gorm.DB) error {
	return db.Delete(r).Error
}

func (r *Room) FindByID(db *gorm.DB, id uint) error {
	defer r.normalize()
	return db.Where("id = ?", id).Find(r).Error
}

func (r *Room) FindByUserID(db *gorm.DB, id uint) error {
	defer r.normalize()
	return db.Joins("INNER JOIN players ON players.room_id = rooms.id").Where("players.user_id = ?", id).Find(r).Error
}

func FindRoomsByGame(db *gorm.DB, game string, bound RowBound) ([]*Room, error) {
	rooms := make([]*Room, 0)
	err := db.Where("game_id = ?", game).Limit(bound.Limit).Offset(bound.Offset).Find(&rooms).Error
	for _, room := range rooms {
		room.normalize()
	}
	return rooms, err
}

func (r *Room) CreateByHost(db *gorm.DB, host *Player) error {
	defer r.normalize()
	host.State = HOST
	host.Seat = 0
	host.normalize()

	r.CreatedTime = time.Now()
	r.Players = []*Player{
		host,
	}
	return db.Save(r).Error
}

func (r *Room) Join(db *gorm.DB, p *Player) error {
	defer r.normalize()
	p.State = NOT_READY
	if seat, ok := r.GetNextSeat(); ok {
		p.Seat = seat
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "The room is full")
	}
	p.normalize()

	r.Players = append(r.Players, p)
	return db.Save(r).Error
}

func (r *Room) Exit(db *gorm.DB, p *Player) error {
	host := p.State == HOST
	r.RemovePlayer(p.UserID)

	tr := db.Begin()
	// update this player
	err := tr.Model(p).Updates(map[string]interface{}{
		"room_id": 0,
		"state":   OUT_OF_GAME,
		"seat":    0,
	}).Error
	if err != nil {
		tr.Rollback()
		return err
	}
	// delete room if no player
	if len(r.Players) == 0 {
		err := tr.Delete(r).Error
		if err != nil {
			tr.Rollback()
			return err
		}
	} else
	// new host if exit player is host
	if host {
		newHostPlayer := r.Players[0]
		err := tr.Model(newHostPlayer).Update("state", HOST).Error
		if err != nil {
			tr.Rollback()
			return err
		}
	}
	tr.Commit()
	return nil
}

func (r *Room) SwapSeat(db *gorm.DB, a, b uint) error {
	xgo.MustTrue(a < r.PlayerCount, "Seat")
	xgo.MustTrue(b < r.PlayerCount, "Seat")

	p, ok := r.FindPlayerBySeat(a)
	if ok {
		p.Seat = b
	}
	p, ok = r.FindPlayerBySeat(b)
	if ok {
		p.Seat = a
	}
	return r.save(db)
}
