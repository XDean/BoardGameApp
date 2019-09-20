package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
	"net/http"
	"strconv"
)

func DBNotFound(err error, msg string) error {
	if gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound, msg)
	} else {
		return err
	}
}

func IntParam(c echo.Context, name string) int {
	param := c.Param(name)
	if i, err := strconv.Atoi(param); err == nil {
		return i
	} else {
		xecho.MustNoError(echo.NewHTTPError(http.StatusBadRequest, "Unrecognized param '"+name+"': "+param))
		return 0
	}
}
