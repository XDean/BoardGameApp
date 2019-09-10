package config

import (
	"github.com/xdean/goex/xconfig"
	"github.com/xdean/goex/xgo"
	wechatConfig "github.com/xdean/miniboardgame/go/wechat/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Debug bool
var SettingFile string
var SecretKey string

var Instance Config

type Config struct {
	DB     DB
	Wechat Wechat
}

type DB struct {
	Dialect string
	URL     string
}

type Wechat struct {
	wechatConfig.Wechat
	AuthUrl string
}

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

func (c *Config) ToYaml() string {
	out, err := yaml.Marshal(c)
	xgo.MustNoError(err)
	return string(out)
}
