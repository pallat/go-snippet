package handler

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Func is our custom Handler function
type Func func(*Context)

// Serve use to convert Func to gin.HandlerFunc
func Serve(h Func, l Logger) gin.HandlerFunc {
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

// Logger represents the logging.
type Logger interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
}
