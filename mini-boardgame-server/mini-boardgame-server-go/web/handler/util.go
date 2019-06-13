package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DBNotFound(err error, msg string) error {
	if gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, msg)
	} else {
		return err
	}
}
