package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func main() {
	r := setupRouter()
	// Listen and Serve on 0.0.0.0:8080
	r.Run(":8080")
}
