package main

import (
	"fmt"
	"project4/controller"
	"project4/infra/postgres"
	"project4/middleware"
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

	transactionHistoryRepository := repository.NewTransactionHistoryRepository(db)
	transactionHistoryService := service.NewTransactionHistoryService(transactionHistoryRepository, productRepository, userRepository)
	transactionHistoryController := controller.NewTransactionHistoryController(transactionHistoryService, productService)

	router := gin.Default()

	// Users
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.PUT("/users/:id", middleware.AuthMiddleware(), userController.UpdateUser)
	router.DELETE("/users/:id", middleware.AuthMiddleware(), userController.DeleteUser)

	// Category
	router.POST("/categories", middleware.AuthMiddleware(), categoryController.CreateCategory)
	router.GET("/categories", middleware.AuthMiddleware(), categoryController.GetCategory)
	router.PUT("/categories/:id", middleware.AuthMiddleware(), categoryController.UpdateCategory)
	router.DELETE("/categories/:id", middleware.AuthMiddleware(), categoryController.DeleteCategory)

	// Product
	router.POST("/products", middleware.AuthMiddleware(), productController.CreateProduct)
	router.GET("/products", middleware.AuthMiddleware(), productController.GetProduct)
	router.PUT("/products/:id", middleware.AuthMiddleware(), productController.UpdateProduct)
	router.DELETE("/products/:id", middleware.AuthMiddleware(), productController.DeleteProduct)

	// Transaction History
	router.POST("/transactionhistories", middleware.AuthMiddleware(), transactionHistoryController.CreateTransactionHistory)
	router.GET("/transactionhistories/my-transactions", middleware.AuthMiddleware(), transactionHistoryController.GetMyTransactionHistory)
	router.GET("/transactionhistories/user-transactions", middleware.AuthMiddleware(), transactionHistoryController.GetAllTransactionHistory)

	router.Run()
}
