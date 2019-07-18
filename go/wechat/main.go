package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/wechat/handler"
	"net/http"
	"os"
)
import "github.com/xdean/miniboardgame/go/wechat/config"

func main() {
	flag.StringVar(&config.SettingFile, "setting", "", "setting file path")
	flag.StringVar(&config.SettingFile, "secret-key", "123456", "secret key")

	flag.Parse()
	if config.SettingFile == "" {
		fmt.Println("Please specify setting file")
		os.Exit(1)
	}
	err := config.Instance.Load(config.SettingFile)
	if err != nil {
		fmt.Println("Fail to load setting file: ", err.Error())
		os.Exit(1)
	}

	e := echo.New()
	e.Validator = xecho.NewValidator()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(xecho.BreakErrorRecover())

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	e.GET("/wechat", handler.CheckSignature)
	e.POST("/wechat", handler.Message)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Instance.Web.Port)))
}
