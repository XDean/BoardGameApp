package log

import (
	"github.com/sirupsen/logrus"
)

type GormLogrusLogger struct {
	Logger *logrus.Logger
}

func (l *GormLogrusLogger) Print(values ...interface{}) {
	entry := l.Logger.WithField("name", "gorm")
	if len(values) > 1 {
		level := values[0]
		source := values[1]
		entry = entry.WithField("source", source)
		if level == "sql" {
			duration := values[2]
			// sql
			entry.WithField("took", duration).Debug(formatGormSql(values))
		} else {
			entry.Error(values[2:]...)
		}
	} else {
		entry.Error(values...)
	}
}
