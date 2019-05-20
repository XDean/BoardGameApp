package web

import (
	"github.com/XDean/MiniBoardgame/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() {
	echo := echo.New()
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	api := echo.Group("/api")

	authored := api.Group("")
	authored.Use(middleware.JWTWithConfig(auth.JwtAuthenticateConfig()))

	admin := authored.Group("")
	admin.Use(auth.AdminAuth)
}
