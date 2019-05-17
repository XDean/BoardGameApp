package log

import (
	"database/sql/driver"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"time"
)

var sqlRegexp = regexp.MustCompile(`(\$\d+)|\?`)

type GormLogger struct {
	Name   string
	Logger *logrus.Logger
}

func (l *GormLogger) Print(values ...interface{}) {
	entry := l.Logger.WithField("name", l.Name)
	if len(values) > 1 {
		level := values[0]
		source := values[1]
		entry = l.Logger.WithField("source", source)
		if level == "sql" {
			duration := values[2]
			// sql
			var formattedValues []interface{}
			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format(time.RFC3339)))
					} else if b, ok := value.([]byte); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", string(b)))
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				}
			}
			entry.WithField("took", duration).Debug(fmt.Sprintf(sqlRegexp.ReplaceAllString(values[3].(string), "%v"), formattedValues...))
		} else {
			entry.Error(values[2:]...)
		}
	} else {
		entry.Error(values...)
	}

}
