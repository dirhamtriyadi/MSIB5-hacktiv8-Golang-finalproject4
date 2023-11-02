package main

import (
	"fmt"
	"project4/controller"
	"project4/infra/postgres"
	"project4/model/entity"
	"project4/repository"
	"project4/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db := postgres.InitDB()
	db.AutoMigrate(&entity.User{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	// Users
	router.POST("/users/register", userController.RegisterUser)
	// router.POST("/users/login", userController.LoginUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	router.Run()
}
