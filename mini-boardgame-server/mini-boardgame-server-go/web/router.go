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

	authGroup := e.Group("/auth")
	authGroup.GET("sign-up", handler.SignUp)

	apiGroup := e.Group("/apiGroup")

	authored := apiGroup.Group("")
	authored.Use(middleware.JWTWithConfig(handler.JwtAuthenticateConfig()))

	admin := authored.Group("")
	admin.Use(handler.AdminAuthMiddleware)
}
