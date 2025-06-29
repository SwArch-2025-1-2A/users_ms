package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/SwArch-2025-1-2A/users_ms/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"username"`
}

func CreateUserHandler(c *gin.Context) {
	app, ok := GetApp(c)
	// The error response and logging is already done by the GetApp function, so we don't need to do it here
	if !ok {
		return
	}

	var requestBody UserResponse
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Expected two parameters in body: id (a valid UUID) and username"})
		log.Println(err.Error())
		return
	}

	args := repository.CreateUserParams{
		ID:         requestBody.ID,
		Name:       requestBody.Name,
		ProfilePic: nil,
	}

	user, err := app.Queries.CreateUser(app.Context, args)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error when creating the user"})
		log.Println(err.Error())
		return
	}

	response := UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"status": "success", "data": response})
}

type GetUserResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"username"`
	ProfilePicUrl string    `json:"profilePicUrl"`
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

	response := GetUserResponse{
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

	response := GetUserResponse{
		ID:            user.ID,
		Name:          user.Name,
		ProfilePicUrl: GenerateImageURL(user.ID),
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "success", "data": response})
}
