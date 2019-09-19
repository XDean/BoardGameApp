package web

import (
	"github.com/labstack/echo/v4"
	_const "github.com/xdean/miniboardgame/go/server/const"
	"github.com/xdean/miniboardgame/go/server/web/handler"
	"github.com/xdean/miniboardgame/go/server/web/middleware"
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

	authored.GET("/player", handler.GetPlayer)
	authored.GET("/player/:id", handler.GetPlayerByID)

	authored.POST("/room", handler.CreateRoom)
	authored.GET("/room/:id", handler.GetRoomByID)
	authored.GET("/room", handler.GetRoom, middleware.AuthRoom())
	authored.GET("/room/join/:id", handler.JoinRoom)
	authored.GET("/room/exit", handler.ExitRoom)

	authored.GET("/games", handler.GetGameList)
	authored.GET("/game/start", handler.StartGame, middleware.AuthRoom())
	authored.POST("/game/event", handler.GameEvent, middleware.AuthRoom())

	authored.GET("/socket/room", handler.RoomSocket, middleware.AuthRoom())
	authored.GET("/sse/room", handler.RoomSSE, middleware.AuthRoom())

	admin := authored.Group("/admin")
	admin.Use(middleware.AuthRole(_const.ROLE_ADMIN))
}
