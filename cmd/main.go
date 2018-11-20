package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"ktb.co.th/api/prototype/api"
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
		c.Logger().Infof("json", map[string]string{
			"state":   "input",
			"id":      c.Response().Header().Get(echo.HeaderXRequestID),
			"payload": string(reqBody),
		})
		c.Logger().Infof("json", map[string]string{
			"state":    "output",
			"id":       c.Response().Header().Get(echo.HeaderXRequestID),
			"response": string(resBody),
		})
	}))

	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader(`{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}`)

	}

	e.Use(logs.MiddlewareWriter(logout))

	// Routes
	e.GET("/", hello)
	e.GET("/bin", api.HTTPBin)
	e.GET("/business/error", api.HTTPBinBusinessFail)
	e.GET("/technical/error", api.HTTPBinTechnicalFail)
	e.GET("/panic", func(c echo.Context) error {
		x := []int{}
		return c.JSON(http.StatusOK, map[string]int{"result": x[10]})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
