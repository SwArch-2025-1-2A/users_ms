package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", pingHandler)

	router.Run("localhost:8008")
}

func pingHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "PONG")
}
