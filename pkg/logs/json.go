package logs

import "github.com/labstack/echo"

const format = "json"

func Success(c echo.Context) (string, map[string]string) {
	return format, map[string]string{
		"service": "HTTPBin",
		"state":   "response",
		"status":  "S",
		"id":      c.Response().Header().Get(echo.HeaderXRequestID),
	}
}

func Info(c echo.Context, msg string) (string, map[string]string) {
	return format, map[string]string{
		"service": "HTTPBin",
		"state":   "response",
		"status":  "S",
		"message": " msg",
		"id":      c.Response().Header().Get(echo.HeaderXRequestID),
	}
}

func Error(c echo.Context, err error) (string, map[string]string) {
	return format, map[string]string{
		"service": "HTTPBin",
		"state":   "request",
		"status":  "E",
		"message": err.Error(),
		"id":      c.Response().Header().Get(echo.HeaderXRequestID),
	}
}
