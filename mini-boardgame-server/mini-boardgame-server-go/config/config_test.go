package config

import (
	"path"
	"runtime"
	"testing"

	"gopkg.in/d4l3k/messagediff.v1"
)

var actual *Config = new(Config)
var expected = Config{
	Debug: true,
	DB: DB{
		Dialect: "sqlite3",
		URL:     "some-url",
	},
}

func TestInit(t *testing.T) {
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("fail to get caller")
	}
	dir := path.Dir(filePath)
	confPath := path.Join(dir, "config_test.yaml")
	err := actual.Load(confPath)
	if err != nil {
		t.Fatal("fail to init config, reason:", err)
	}
	if diff, equal := messagediff.PrettyDiff(expected, *actual); !equal {
		t.Errorf("expected Global = %#v\n%s", expected, diff)
	}
}
