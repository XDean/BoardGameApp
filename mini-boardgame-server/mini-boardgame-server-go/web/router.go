package web

import (
	"github.com/XDean/MiniBoardgame/config"
	"github.com/XDean/MiniBoardgame/handler"
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

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	authGroup := e.Group("/auth")
	authGroup.GET("sign-up", handler.SignUp)

	apiGroup := e.Group("/apiGroup")

	authored := apiGroup.Group("")
	authored.Use(middleware.JWTWithConfig(JwtAuthenticateConfig()))

	admin := authored.Group("")
	admin.Use(AdminAuth)
}
