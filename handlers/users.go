package handlers

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/SwArch-2025-1-2A/users_ms/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	ProfilePicUrl string    `json:"profilePicUrl"`
}

func CreateUserHandler(c *gin.Context) {
	app, ok := GetApp(c)
	if !ok {
		return
	}

	idStr := c.PostForm("id")
	if idStr == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "You didn't enter an id (uuid)"})
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "The given id is not a valid UUID"})
		return
	}
	name := c.PostForm("name")

	file, header, err := c.Request.FormFile("profilePic")
	var profilePic []byte
	// A missing profilePic shouldn't cause any errors. It should be possible
	// to create a user without a profilePic and leave that for later
	if err == http.ErrMissingFile {
		profilePic = nil
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Strange error while reading the profile pic"})
		log.Println(err.Error())
		return
	} else {
		defer file.Close()
		if !strings.HasPrefix(header.Header.Get("Content-Type"), "image/") {
			c.IndentedJSON(http.StatusBadRequest,
				gin.H{"error": "The attached file isn't an image. Make sure that the content-type header is set to image/"})
			return
		}

		profilePic, err = io.ReadAll(file)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error when trying to read the image"})
			log.Println(err.Error())
			return
		}
	}

	args := repository.CreateUserParams{
		ID:         id,
		Name:       name,
		ProfilePic: profilePic,
	}

	user, err := app.Queries.CreateUser(app.Context, args)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error when creating the user"})
		log.Println(err.Error())
		return
	}

	userResponse := UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		ProfilePicUrl: GenerateImageURL(user.ID),
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"status": "success", "data": userResponse})
}
