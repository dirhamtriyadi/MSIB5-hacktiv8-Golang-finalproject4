package controller

import (
	"fmt"
	"net/http"
	"project4/helper"
	"project4/middleware"
	"project4/model/input"
	"project4/model/response"
	"project4/service"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input input.UserRegisterInput

	err := c.ShouldBindJSON(&input)

	user, err := govalidator.ValidateStruct(input)

	if !user {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err.Error(),
		})
		fmt.Println("error: " + err.Error())
		return
	}

	result, err := h.userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err.Error(),
		})
		fmt.Println("error: " + err.Error())
		return
	}

	registerResponse := response.UserRegisterResponse{
		ID:        result.ID,
		FullName:  result.FullName,
		Email:     result.Email,
		Password:  result.Password,
		Balance:   result.Balance,
		CreatedAt: result.CreatedAt,
	}

	response := helper.APIResponse("created", registerResponse)
	c.JSON(201, response)
}

func (h *userController) LoginUser(c *gin.Context) {
	var input input.UserLoginInput

	err := c.ShouldBindJSON(&input)

	login, err := govalidator.ValidateStruct(input)

	if !login {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})

		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	// send to services
	// get user by email
	user, err := h.userService.GetUserByEmail(input.Email)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// return when user not found!
	if user.ID == 0 {
		errorMessages := "User not found!"
		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		response := helper.APIResponse("failed", "password not match!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create token
	jwtService := middleware.NewService()
	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		response := helper.APIResponse("failed", "failed to generate token!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginResponse := response.UserLoginResponse{
		Token: token,
	}

	// return token
	response := helper.APIResponse("ok", loginResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userController) UpdateUser(c *gin.Context) {
	var inputUserUpdate input.UserUpdateInput

	err := c.ShouldBindJSON(&inputUserUpdate)

	user, err := govalidator.ValidateStruct(inputUserUpdate)

	if !user {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	var idUserUri input.UserUpdateID

	err = c.ShouldBindUri(&idUserUri)

	userId, err := govalidator.ValidateStruct(idUserUri)

	if !userId {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	id_user := idUserUri.ID

	_, err = h.userService.UpdateUser(id_user, inputUserUpdate)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	userUpdated, err := h.userService.GetUserByID(id_user)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	updatedResponse := response.UserUpdateResponse{
		ID:        userUpdated.ID,
		FullName:  userUpdated.FullName,
		Email:     userUpdated.Email,
		Password:  userUpdated.Password,
		Balance:   userUpdated.Balance,
		UpdatedAt: userUpdated.UpdatedAt,
	}

	response := helper.APIResponse("ok", updatedResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userController) DeleteUser(c *gin.Context) {
	var idUserUri input.UserDeleteID

	err := c.ShouldBindUri(&idUserUri)

	userId, err := govalidator.ValidateStruct(idUserUri)

	if !userId {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	id_user := idUserUri.ID

	_, err = h.userService.DeleteUser(id_user)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	deletedResponse := response.UserDeleteResponse{
		Message: "User has been deleted",
	}

	response := helper.APIResponse("ok", deletedResponse)
	c.JSON(http.StatusOK, response)
}
