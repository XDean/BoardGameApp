package web

import (
	_const "github.com/XDean/MiniBoardgame/const"
	"github.com/XDean/MiniBoardgame/web/handler"
	"github.com/XDean/MiniBoardgame/web/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRouter(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	loginGroup := e.Group("/auth")
	loginGroup.GET("/signup", handler.SignUp)
	loginGroup.GET("/login", handler.Login)
	loginGroup.GET("/logout", handler.Logout, middleware.Authorized())

	apiGroup := e.Group("/api")

	authored := apiGroup.Group("")
	authored.Use(middleware.Authorized())
	authored.GET("/user", handler.GetUser)
	authored.GET("/user/:id", handler.GetUserById)
	authored.GET("/user/profile", handler.GetProfile)
	authored.GET("/user/profile/:id", handler.GetProfileById)
	authored.PATCH("/user/profile", handler.UpdateProfile)

	admin := authored.Group("")
	admin.Use(middleware.AuthRole(_const.ROLE_ADMIN))
}
