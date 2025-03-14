package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExampleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to MMS - Message Management System",
		"status":  "success",
	})
}
