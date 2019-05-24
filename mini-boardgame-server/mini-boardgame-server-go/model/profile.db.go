package model

import "github.com/jinzhu/gorm"

func (profile *Profile) Save(db *gorm.DB) error {
	return db.Save(profile).Error
}

func (profile *Profile) FindByUserID(db *gorm.DB, id uint) error {
	if err := db.Where("user_id = ?", id).Find(profile).Error; err != nil {
		return err
	}
	return nil
}
