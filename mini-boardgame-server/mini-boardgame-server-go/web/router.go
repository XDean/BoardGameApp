package web

import (
	"github.com/XDean/MiniBoardgame/config"
	"github.com/XDean/MiniBoardgame/handler"
	"github.com/XDean/MiniBoardgame/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func InitRouter() {
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

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	loginGroup := e.Group("/auth")
	loginGroup.GET("sign-up", handler.SignUp)
	loginGroup.GET("login", handler.Login)

	apiGroup := e.Group("/apiGroup")

	authored := apiGroup.Group("")
	authored.Use(handler.AuthMiddleware())

	admin := authored.Group("")
	admin.Use(handler.AdminAuthMiddleware)

}
