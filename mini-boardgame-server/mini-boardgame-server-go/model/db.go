package model

import (
	"github.com/jinzhu/gorm"

	"github.com/XDean/MiniBoardgame/config"
	"github.com/XDean/MiniBoardgame/log"
	// load mysql Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// load sqlite Driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func LoadFromConfig() (*gorm.DB, error) {
	db, err := gorm.Open(config.Global.DB.Dialect, config.Global.DB.URL)
	if err != nil {
		return nil, err
	}
	return db, ConfigDB(db)
}

func ConfigDB(database *gorm.DB) error {
	database.SetLogger(&log.GormLogger{
		Name:   "DB",
		Logger: log.Global,
	})

	// foreign key constraint is disable in SQLite by default, should enable it first
	err := database.Exec("PRAGMA foreign_keys=ON;").Error
	if err != nil {
		return err
	}

	// Db.ShowSQL(true)
	database.LogMode(config.Global.Debug)
	database = database.Set("gorm:auto_preload", true)

	err = database.AutoMigrate(new(User), new(Role), new(Profile), new(Room), new(Player)).Error
	return err
}
