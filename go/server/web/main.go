package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/server/config"
	"github.com/xdean/miniboardgame/go/server/log"
	"github.com/xdean/miniboardgame/go/server/model"
	"github.com/xdean/miniboardgame/go/server/web/handler"
	myMiddleware "github.com/xdean/miniboardgame/go/server/web/middleware"
	"net/http"
)

func Run() {
	// Load Config
	err := config.Instance.Load("./config-dev.yml")
	if err != nil {
		log.Global.Fatal("Config load fail", err)
	}
	// Load DB
	db, err := model.LoadFromConfig()
	if err != nil {
		log.Global.Fatal("Database can't be loaded from config", err)
	}

	// Init echo
	e := echo.New()

	e.Validator = handler.NewValidator()

	e.HTTPErrorHandler = errorHandler

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(xecho.BreakErrorRecover())
	e.Use(myMiddleware.DbContext(db))
	e.Use(myMiddleware.Jwt())

	InitRouter(e)

	log.Global.Fatal(e.Start(":80"))
}

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := interface{}(nil)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if he.Internal == nil {
			msg = xecho.J{
				"message": he.Message,
			}
		} else {
			msg = xecho.J{
				"message": he.Message,
				"cause":   he.Internal.Error(),
			}
		}
	} else {
		msg = xecho.J{
			"message": http.StatusText(code),
			"cause":   err.Error(),
		}
	}
	if _, ok := msg.(string); ok {
		msg = echo.Map{"message": msg}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			log.Global.Error(err)
		}
	}
}
