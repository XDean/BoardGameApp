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
	defer r.Normalize()
	return db.Where("id = ?", id).Find(r).Error
}

func FindRoomsByGame(db *gorm.DB, game string, bound RowBound) ([]*Room, error) {
	rooms := make([]*Room, 0)
	err := db.Where("game_name = ?", game).Limit(bound.Limit).Offset(bound.Offset).Find(&rooms).Error
	for _, room := range rooms {
		room.Normalize()
	}
	return rooms, err
}

func (r *Room) CreateByHost(db *gorm.DB, host *User) error {
	defer r.Normalize()
	r.CreatedTime = time.Now()
	r.Players = []*Player{
		{
			UserID: host.ID,
			State:  HOST,
		},
	}
	return db.Save(r).Error
}
