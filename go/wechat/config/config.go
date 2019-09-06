package config

import (
	"github.com/xdean/goex/xconfig"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Debug bool
var SettingFile string
var SecretKey string

var Instance Config

// Conf is the root configuration struct
type (
	Config struct {
		Web    Web
		Wechat Wechat
	}

	Web struct {
		Port int
	}

	Wechat struct {
		Url       string
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
