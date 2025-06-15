package handlers

import (
	"database/sql"
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
	// The error response and logging is already done by the GetApp function, so we don't need to do it here
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

func GetUserHandler(c *gin.Context) {
	app, ok := GetApp(c)
	// The error response and logging is already done by the GetApp function, so we don't need to do it here
	if !ok {
		return
	}

	id, ok := ReadIdParam(c)
	if !ok {
		return
	}

	user, err := app.Queries.GetUserById(app.Context, id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No user with that id in the database"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error when trying to retrieve the user from the DB"})
		log.Println(err.Error())
		return
	}

	response := UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		ProfilePicUrl: GenerateImageURL(user.ID),
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

type ChangeUsernameRequestBody struct {
	Name string `json:"name"`
}

func ChangeUserName(c *gin.Context) {
	app, ok := GetApp(c)
	// The error response and logging is already done by the GetApp function, so we don't need to do it here
	if !ok {
		return
	}

	id, ok := ReadIdParam(c)
	if !ok {
		return
	}

	var requestBody ChangeUsernameRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "You have to provide a name (and only a name) in the request body"})
		// I am not sure that a bad request is the only reason why this error might happen, so I want to log them to check
		log.Println(err.Error())
		return
	}

	args := repository.ChangeUserNameParams{
		ID:   id,
		Name: requestBody.Name,
	}
	user, err := app.Queries.ChangeUserName(app.Context, args)
	if err != nil {
		// I have no easy way to know what went wrong, so I assume an internal server error
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error when trying to update the username"})
		log.Println(err.Error())
		return
	}

	response := UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		ProfilePicUrl: GenerateImageURL(user.ID),
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "success", "data": response})
}
