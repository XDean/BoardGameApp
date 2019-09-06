package model

import (
	"github.com/jinzhu/gorm"

	"github.com/xdean/miniboardgame/go/server/config"
	"github.com/xdean/miniboardgame/go/server/log"
	// load mysql Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// load sqlite Driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func LoadFromConfig() (*gorm.DB, error) {
	db, err := gorm.Open(config.Instance.DB.Dialect, config.Instance.DB.URL)
	if err != nil {
		return nil, err
	}
	return Configure(db)
}

func Configure(database *gorm.DB) (*gorm.DB, error) {
	database.SetLogger(&log.GormLogrusLogger{
		Logger: log.Global,
	})

	// foreign key constraint is disable in SQLite by default, should enable it first
	err := database.Exec("PRAGMA foreign_keys=ON;").Error
	if err != nil {
		return nil, err
	}

	// Db.ShowSQL(true)
	database = database.LogMode(config.Debug)
	database = database.Set("gorm:auto_preload", true)

	err = database.AutoMigrate(new(User), new(Role), new(Profile), new(Room), new(Player)).Error
	return database, err
}
