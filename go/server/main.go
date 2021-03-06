package main

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/xdean/miniboardgame/go/server/config"
	"github.com/xdean/miniboardgame/go/server/web"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "Wechat BG Mini Program"
	app.Usage = "Run XDean Wechat BG Mini Program Server"

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
		fmt.Println("Config", config.Instance.ToYaml())
		web.Run()
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
