package middleware

import (
	"github.com/XDean/MiniBoardgame/model"
	"github.com/labstack/echo/v4"
)

func BreakErrorRecover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					e, ok := r.(model.BreakError)
					if !ok {
						panic(r)
					}
					err = e.Actual
				}
			}()
			return next(c)
		}
	}
}
