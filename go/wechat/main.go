package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/xdean/goex/xecho"
	"github.com/xdean/miniboardgame/go/wechat/config"
	"github.com/xdean/miniboardgame/go/wechat/handler"
	"github.com/xdean/miniboardgame/go/wechat/service"
	"net/http"
	"os"
)

func main() {

	app := cli.NewApp()

	app.Name = "Wechat BG"
	app.Usage = "Run XDean Wechat BG Server"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "setting,s",
			Usage:       "Setting file path",
			Destination: &config.SettingFile,
		},
		cli.StringFlag{
			Name:        "key,k",
			Usage:       "Secret key",
			Destination: &config.SecretKey,
		},
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "Debug mode",
			Destination: &config.Debug,
		},
	}

	app.Action = func(c *cli.Context) error {
		if config.SettingFile == "" {
			return errors.New("Please specify setting file")
		}
		err := config.Instance.Load(config.SettingFile)
		if err != nil {
			return err
		}
		if config.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		fmt.Println(config.Instance)
		service.StartAccessTokenTask()
		run()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
func run() {
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
