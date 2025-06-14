package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", pingHandler)

	router.Run()
}

func pingHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "PONG")
}
