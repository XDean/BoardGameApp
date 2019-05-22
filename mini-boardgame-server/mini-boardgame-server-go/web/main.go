package web

import (
	"github.com/XDean/MiniBoardgame/config"
	"github.com/XDean/MiniBoardgame/handler"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Run() {
	config.Global.Load("./config-dev.yml")

	e := echo.New()
	e.Validator = handler.NewValidator()

	if config.Global.Debug {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	db, err := model.LoadFromConfig()
	if err != nil {
		panic("Database can't be loaded from config: " + err.Error())
	}
	e.Use(handler.DatabaseContextMiddleware(db))

	InitRouter(e)

	log.Fatal(e.Start(":80"))
}
