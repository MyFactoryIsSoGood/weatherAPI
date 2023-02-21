package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()
	app.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "true"})
	})
}
