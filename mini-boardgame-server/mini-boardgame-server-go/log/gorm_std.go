package log

import (
	"fmt"
)

type GormStdLogger struct {
}

func (t GormStdLogger) Print(values ...interface{}) {
	if len(values) > 1 {
		level := values[0]
		source := values[1]
		if level == "sql" {
			duration := values[2]
			// sql
			sql := formatGormSql(values)
			fmt.Println(sql,
				"AFFECT:", values[5], "DURATION:", duration, source)
		} else {
			fmt.Println(values[2:]...)
		}
	} else {
		fmt.Println(values...)
	}
}
