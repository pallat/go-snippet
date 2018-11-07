package api

import (
	"github.com/gin-gonic/gin"
	"ktb.co.th/api/prototype/pkg/handler"
)

func Feature(c *handler.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
