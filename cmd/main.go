package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"ktb.co.th/api/prototype/pkg/logs"
)

// Handler
func hello(c echo.Context) error {
	c.Logger().Info(c.Response().Header().Get(echo.HeaderXRequestID))

	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	var logout io.Writer = os.Stdout

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Skipper: middleware.DefaultSkipper,
			Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
				`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
				`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
				`"bytes_out":${bytes_out}}` + "\n",
			Output: logout,
		},
	))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		c.Logger().SetOutput(logout)
		c.Logger().Info(c.Response().Header().Get(echo.HeaderXRequestID) + "," + string(reqBody))
		c.Logger().Info(c.Response().Header().Get(echo.HeaderXRequestID) + "," + string(resBody))
	}))

	e.Use(logs.MiddlewareWriter(logout))

	// Routes
	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":1323"))
}
