package handler

import (
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
	formatter:= user.FormatUser(newUser, "this is token")

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