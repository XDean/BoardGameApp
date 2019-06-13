package handler

import (
	"github.com/dlclark/regexp2"
	"github.com/labstack/echo/v4"
	"github.com/xdean/goex/xecho"
)

var Regexps = map[string]*regexp2.Regexp{
	"USERNAME": regexp2.MustCompile("^(?!_)(?![0-9]+$)[a-zA-Z0-9_]+(?<!_)$", 0),
	"PASSWORD": regexp2.MustCompile("^(?![0-9]+$)(?![a-zA-Z]+$)(?=[A-Za-z])[0-9A-Za-z]{6,16}$", 0),
}

func NewValidator() echo.Validator {
	xecho.ValidateRegexps = Regexps
	return xecho.NewValidator()
}
