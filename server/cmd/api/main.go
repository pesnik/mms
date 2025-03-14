package main

import (
	"github.com/pesnik/mms/internal/handlers"
	"github.com/pesnik/mms/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	loggerInit()
	r := gin.Default()
	r.GET("/", handlers.ExampleHandler)
	r.Run(":8080")
}

func loggerInit() {
	logger.Log = logger.New()
}
