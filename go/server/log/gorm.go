package log

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"time"
)

var sqlRegexp = regexp.MustCompile(`(\$\d+)|\?`)

func formatGormSql(values []interface{}) string {
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
	sql := fmt.Sprintf(sqlRegexp.ReplaceAllString(values[3].(string), "%v"), formattedValues...)
	return sql
}
