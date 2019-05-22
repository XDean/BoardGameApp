package web

import (
	"github.com/XDean/MiniBoardgame/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRouter(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		if user, err := handler.GetCurrentUser(c); err == nil {
			return c.JSON(http.StatusOK, "hello "+user.Username)
		} else {
			return c.JSON(http.StatusOK, "pong")
		}
	})

	loginGroup := e.Group("/auth")
	loginGroup.GET("/sign-up", handler.SignUp)
	loginGroup.GET("/login", handler.Login)

	apiGroup := e.Group("/apiGroup")

	authored := apiGroup.Group("")
	authored.Use(handler.AuthMiddleware())

	admin := authored.Group("")
	admin.Use(handler.AdminAuthMiddleware)
}
