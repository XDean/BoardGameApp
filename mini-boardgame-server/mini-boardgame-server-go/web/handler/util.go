package handler

import (
	"github.com/XDean/MiniBoardgame/model"
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

func BindAndValidate(c echo.Context, param interface{}) {
	if err := c.Bind(param); err != nil {
		panic(model.BreakError{Actual: err})
	}
	if err := c.Validate(param); err != nil {
		panic(model.BreakError{Actual: err})
	}
}

func MustNoError(err error) {
	if err != nil {
		panic(model.BreakError{Actual: err})
	}
}
