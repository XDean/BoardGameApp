package config

import (
	"flag"
	"github.com/xdean/goex/xconfig"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var SecretKey = flag.String("secret-key", "123456", "secret key")

var Instance Config

func init() {
	flag.Parse()
}

// Conf is the root configuration struct
type Config struct {
	Debug  bool
	Wechat Wechat
}

type Wechat struct {
	AppId     string
	AppSecret string
	AuthUrl   string
}

// Init configuration instance
func (c *Config) Load(path string) (err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, c)
	if err == nil {
		err = xconfig.Decode(c, *SecretKey)
	}
	return
}
