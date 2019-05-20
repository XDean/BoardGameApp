package db

import (
	"errors"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/jinzhu/gorm"

	"github.com/XDean/MiniBoardgame/config"
	"github.com/XDean/MiniBoardgame/log"
	// load mysql Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// load sqlite Driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB instance
var DB *gorm.DB

func initDB() error {
	if DB != nil {
		return errors.New("DB has been inited")
	}
	var err error
	DB, err = gorm.Open(config.Global.DB.Dialect, config.Global.DB.URL)
	if err != nil {
		return err
	}
	DB.SetLogger(&log.GormLogger{
		Name:   "DB",
		Logger: log.Global,
	})

	// foreign key constraint is disable in SQLite by default, should enable it first
	err = DB.Exec("PRAGMA foreign_keys=ON;").Error
	if err != nil {
		return err
	}

	// Db.ShowSQL(true)
	DB.LogMode(config.Global.Debug)
	DB = DB.Set("gorm:auto_preload", true)

	err = DB.AutoMigrate(new(model.User), new(model.Role), new(model.Profile), new(model.Room), new(model.Player)).Error
	return err
}
