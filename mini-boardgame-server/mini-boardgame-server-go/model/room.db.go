package model

import (
	"github.com/jinzhu/gorm"
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
	err := db.Where("game_name = ?", game).Limit(bound.Limit).Offset(bound.Offset).Find(&rooms).Error
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
