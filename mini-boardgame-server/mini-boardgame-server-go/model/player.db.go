package model

import "github.com/jinzhu/gorm"

func (p *Player) GetByUserID(db *gorm.DB, id uint) error {
	p.UserID = id
	err := db.FirstOrCreate(p, "user_id = ?", id).Error
	if err == nil {
		err = p.normalizeDB(db)
	}
	return err
}

func (p *Player) save(db *gorm.DB) error {
	return db.Save(p).Error
}

func (p *Player) normalizeDB(db *gorm.DB) error {
	p.normalize()
	if p.RoomID != 0 {
		room := new(Room)
		if err := room.FindByID(db, p.RoomID); err != nil {
			return err
		}
		for i, player := range room.Players {
			if player.UserID == p.UserID {
				room.Players[i] = p
				break
			}
		}
		room.normalize()
	}
	return nil
}
