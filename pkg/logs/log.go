package logs

import (
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func MiddlewareWriter(w io.Writer) echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Logger().SetOutput(w)
			c.Logger().SetLevel(log.INFO)
			return handler(c)
		}
	}
}
