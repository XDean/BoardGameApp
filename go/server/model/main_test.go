package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xdean/miniboardgame/go/server/config"
	"github.com/xdean/miniboardgame/go/server/log"
	"io/ioutil"
	"os"
	"testing"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	config.Debug = true
	tmp, err := ioutil.TempFile("", "mini-bg-*.db")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmp.Name())
	fmt.Println("Temp database file:", tmp.Name())
	db, err := gorm.Open("sqlite3", tmp.Name())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db, err = Configure(db)
	if err != nil {
		panic(err)
	}
	db.SetLogger(log.GormStdLogger{})
	testDB = db

	result := m.Run()

	os.Exit(result)
}
