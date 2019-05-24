package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

type J map[string]interface{}

func M(msg string) interface{} {
	return J{
		"message": msg,
	}
}

func DBNotFound(err error, msg string) error {
	if gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, msg)
	} else {
		return err
	}
}
