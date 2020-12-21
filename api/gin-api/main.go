package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/api", func(cnt *gin.Context) {
		cnt.JSON(http.StatusOK, gin.H{
			"message": "Hello world, Gin Golang API server",
		})
	})
	engine.Run(":1389")
}
