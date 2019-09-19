package log

import (
	"github.com/sirupsen/logrus"
)

var Global *logrus.Logger = logrus.StandardLogger()

func init() {
	Global.SetFormatter(&logrus.TextFormatter{})
}
