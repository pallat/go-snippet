package handler

import "github.com/gin-gonic/gin"

type Context struct {
	id string
	*gin.Context
	Logger
}
