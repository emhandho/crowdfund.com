package handler

import (
	"fmt"
	"log"
	"net/http"

	"crowdfund.com/helper"
	"crowdfund.com/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//Catch input from user
	//Input mapping from user to struct registerUserInput
	//passing above struct as parameter to Service
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		res := helper.APIResponse("Register account failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		res := helper.APIResponse("Register account failed.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// token, err := h.jwtService.GenerateToken()
	formatter := user.FormatUser(newUser, "this is token")

	res := helper.APIResponse("Account has been registered.", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, res)
}

func (h *userHandler) Login(c *gin.Context) {
	// User sign input(email & password)
	// Handle catch input
	// mapping from user input to struct input
	// input struct passing to service
	// in service we search with repository help using email x as parameter
	// and then compare the password
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		res := helper.APIResponse("Login failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		res := helper.APIResponse("Login failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	formatter := user.FormatUser(loggedInUser, "this is the token")
	res := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusUnprocessableEntity, res)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// There is an input email from user
	// mapping input email to input struct
	// passing input struct to service
	// service will call repository to determine email has ben registered or not
	// querying repository to database
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		res := helper.APIResponse("Login failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error!"}
		res := helper.APIResponse("Login failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered."

	if isEmailAvailable {
		metaMessage = "Email is availbale."
	}

	res := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// catch input from user
	// Save image to folder "images/"
	// on Service we call repository
	// Who access the Feature, using JWT (but now we use dummy first using ID = 1)
	// repo catch user with ID = 1
	// repo update data user save the file's location

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	log.Println(file.Filename)

	// this userID get from JWT
	userID := 2

	// common path images/filname.jpg
	// path := "images/" + file.Filename

	// best practice path images/userid-filename.jpg
	path := fmt.Sprintf("images/%d-%s", userID,file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfully uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
