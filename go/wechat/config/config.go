package config

import (
	"flag"
	"fmt"
	"github.com/xdean/goex/xconfig"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var SettingFile string
var SecretKey string

var Instance Config

func Init() {
	flag.Parse()
	if SettingFile == "" {
		fmt.Println("Please specify setting file")
		os.Exit(1)
	}
	err := Instance.Load(SettingFile)
	if err != nil {
		fmt.Println("Fail to load setting file: ", err.Error())
		os.Exit(1)
	}
}

// Conf is the root configuration struct
type (
	Config struct {
		Debug  bool
		Web    Web
		Wechat Wechat
	}

	Web struct {
		Port int
	}

	Wechat struct {
		Token     string
		AppId     string
		AppSecret string
		Id        string
		Account   string
	}
)

// Init configuration instance
func (c *Config) Load(path string) (err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, c)
	if err == nil {
		err = xconfig.Decode(c, SecretKey)
	}
	return
}
