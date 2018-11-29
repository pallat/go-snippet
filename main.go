package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pallat/gosnippet/api/example"
	"github.com/pallat/gosnippet/pkg/logs"
)

const (
	maxprocs = 1

	logFormat = `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
		`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
		`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
		`"bytes_out":${bytes_out}}` + "\n"
)

var (
	buildstamp string
	githash    string
)

// Handler
func build(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"buildstamp": buildstamp,
		"githash":    githash,
	})
}

func main() {
	runtime.GOMAXPROCS(maxprocs)

	var logout io.Writer = os.Stdout

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Skipper: middleware.DefaultSkipper,
			Format:  logFormat,
			Output:  logout,
		},
	))

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
	e.GET("/build", build)
	e.GET("/bin", example.HTTPBin)
	e.GET("/business/error", example.HTTPBinBusinessFail)
	e.GET("/technical/error", example.HTTPBinTechnicalFail)
	e.GET("/panic", func(c echo.Context) error {
		x := []int{}
		return c.JSON(http.StatusOK, map[string]int{"result": x[10]})
	})

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
