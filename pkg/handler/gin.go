package handler

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type HandlerFunc func(*Context)

type Context struct {
	id string
	*gin.Context
	Logger
}

func Handler(h HandlerFunc, l Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.NewV4().String()

		ctx := &Context{
			id:      id,
			Context: c,
			Logger:  l,
		}
		h(ctx)
	}
}

type Logger interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
}
