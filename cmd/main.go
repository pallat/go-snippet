package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"ktb.co.th/api/prototype/api"
	"ktb.co.th/api/prototype/pkg/handler"
	"ktb.co.th/api/prototype/pkg/logs"
)

func main() {
	r := gin.Default()
	l := logs.New(os.Stdout)
	r.GET("/ping", handler.Serve(api.Feature, l))
	r.Run() // listen and serve on 0.0.0.0:8080
}
