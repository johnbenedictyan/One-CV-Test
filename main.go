package main

import (
	"net/http"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(logger.SetLogger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
