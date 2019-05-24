package handler

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"reflect"
)

type Validator struct {
	Validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err == nil {
		return nil
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
}

func NewValidator() *Validator {
	v := validator.New()
	_ = v.RegisterValidation("regexp", validRegexp)
	return &Validator{Validator: v}
}

var REGEX = map[string]*regexp2.Regexp{
	"USERNAME": regexp2.MustCompile("^(?!_)(?![0-9]+$)[a-zA-Z0-9_]+(?<!_)$", 0),
	"PASSWORD": regexp2.MustCompile("^(?![0-9]+$)(?![a-zA-Z]+$)(?=[A-Za-z])[0-9A-Za-z]{6,16}$", 0),
}

func validRegexp(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		if str == "" {
			return true
		}
		if regex, find := REGEX[param]; find {
			isMatch, _ := regex.MatchString(str)
			return isMatch
		} else {
			panic(fmt.Sprintf("Bad regex key %v", str))
		}
	default:
		panic(fmt.Sprintf("Bad field type %T", field.Interface()))
	}
}
