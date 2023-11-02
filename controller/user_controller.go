package controller

import (
	"fmt"
	"net/http"
	"project4/helper"
	"project4/model/input"
	"project4/model/response"
	"project4/service"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (h *UserController) RegisterUser(c *gin.Context) {
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
		CreatedAt: result.CreatedAt,
	}

	response := helper.APIResponse("created", registerResponse)
	c.JSON(201, response)
}

func (h *UserController) UpdateUser(c *gin.Context) {
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
		UpdatedAt: userUpdated.UpdatedAt,
	}
	
	response := helper.APIResponse("ok", updatedResponse)
	c.JSON(http.StatusOK, response)
}

func (h *UserController) DeleteUser(c *gin.Context) {
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