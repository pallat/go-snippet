package api

import (
	"ktb.co.th/api/prototype/pkg/handler"
)

func Feature(c *handler.Context) {
	c.JSON(200, map[string]string{
		"message": "pong",
	})
}
