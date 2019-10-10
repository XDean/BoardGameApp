package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xgo"
	"github.com/xdean/miniboardgame/go/games/rps"
	"github.com/xdean/miniboardgame/go/server/config"
	"github.com/xdean/miniboardgame/go/server/log"
	"github.com/xdean/miniboardgame/go/server/model"
	"os"
	"testing"
)

var echoContext *echo.Echo
var dbContext *gorm.DB

func TestMain(m *testing.M) {
	config.Debug = true
	echoContext = echo.New()
	echoContext.Debug = true
	echoContext.Validator = NewValidator()

	result := m.Run()

	os.Exit(result)
}

func resetDB() func() {
	db, err := gorm.Open("sqlite3", ":memory:")
	xgo.MustNoError(err)
	db, err = model.Configure(db)
	xgo.MustNoError(err)
	db.SetLogger(log.GormStdLogger{})
	dbContext = db

	rps.Reset()

	return func() {
		db.Close()
	}
}
