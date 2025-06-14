package main

import (
	"net/http"

	"github.com/SwArch-2025-1-2A/users_ms/app"
	"github.com/SwArch-2025-1-2A/users_ms/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	app := app.NewApp()
	defer app.DBPool.Close()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("app", app)
		c.Next()
	})

	router.GET("/ping", pingHandler)

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("", handlers.CreateUserHandler)
		}
	}

	router.Run()
}

func pingHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "PONG")
}
