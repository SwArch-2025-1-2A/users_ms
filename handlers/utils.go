package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/SwArch-2025-1-2A/users_ms/app"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetApp(c *gin.Context) (*app.App, bool) {
	appInterface, exists := c.Get("app")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("app configuration not available")
		return nil, false
	}

	app, ok := appInterface.(*app.App)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("wrong app type")
		return nil, false
	}

	return app, true
}

// Generate the Image URL
func GenerateImageURL(id uuid.UUID) string {
	port := os.Getenv("PORT")
	hostname := os.Getenv("LOCALHOST")
	return "http://" + hostname + ":" + port + "/api/images/" + id.String()
}
