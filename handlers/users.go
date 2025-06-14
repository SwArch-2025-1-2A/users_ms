package handlers

import "github.com/gin-gonic/gin"

func createUserHandler(c *gin.Context) {
	app, ok := GetApp(c)
	if !ok {
		return
	}

}
