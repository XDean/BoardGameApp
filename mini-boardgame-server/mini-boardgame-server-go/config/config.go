package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Global instance
var Global Config

// Conf is the root configuration struct
type Config struct {
	Debug    bool
	Security Security
	DB       DB
}

type Security struct {
	Key string
}

// DB is Database configuration struct
type DB struct {
	Dialect string
	URL     string
}

// Init configuration instance
func (c *Config) Load(path string) (err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, c)
	return
}
