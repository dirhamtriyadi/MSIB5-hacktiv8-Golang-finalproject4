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
	db.AutoMigrate(&entity.Category{})
	db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&entity.TransactionHistory{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService, productService)

	router := gin.Default()

	// Users
	router.POST("/users/register", userController.RegisterUser)
	// router.POST("/users/login", userController.LoginUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	// Category
	router.POST("/categories", categoryController.CreateCategory)
	router.GET("/categories", categoryController.GetCategory)
	router.PUT("/categories/:id", categoryController.UpdateCategory)
	router.DELETE("/categories/:id", categoryController.DeleteCategory)

	// Product
	router.POST("/products", productController.CreateProduct)
	router.GET("/products", productController.GetProduct)
	router.PUT("/products/:id", productController.UpdateProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)

	router.Run()
}
