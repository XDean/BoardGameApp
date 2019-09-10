package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/config"
	"github.com/xdean/miniboardgame/go/server/model"
	"github.com/xdean/miniboardgame/go/server/web/handler"
	myMiddleware "github.com/xdean/miniboardgame/go/server/web/middleware"
)

func Run() {
	// Load Config
	err := config.Instance.Load("./config-dev.yml")
	if err != nil {
		log.Fatal("Config load fail", err)
	}
	// Load DB
	db, err := model.LoadFromConfig()
	if err != nil {
		log.Fatal("Database can't be loaded from config", err)
	}

	// Init echo
	e := echo.New()

	e.Validator = handler.NewValidator()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(xecho.BreakErrorRecover())
	e.Use(myMiddleware.DbContext(db))
	e.Use(myMiddleware.Jwt())

	InitRouter(e)

	log.Fatal(e.Start(":80"))
}
