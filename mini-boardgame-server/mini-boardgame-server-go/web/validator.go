package web

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewValidator() *Validator {
	v := validator.New()
	v.RegisterValidation("regexp", validRegexp)
	return &Validator{validator: v}
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
		if regex, find := REGEX[param]; find {
			isMatch, _ := regex.MatchString(field.String())
			return isMatch
		} else {
			panic(fmt.Sprintf("Bad regex key %V", field.String()))
		}
	default:
		panic(fmt.Sprintf("Bad field type %T", field.Interface()))
	}
}
