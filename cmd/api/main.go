package main

import (
	"context"
	"log"
	"mini-ecommerce/internal/database"
	"mini-ecommerce/internal/handler"
	"mini-ecommerce/internal/middleware"
	"mini-ecommerce/internal/repository"
	"mini-ecommerce/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load .env: %v", err)
	}

	ctx := context.Background()

	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect db m: %v", err)
	}
	defer db.Close()

	productRepository := repository.NewProductRepositoryImpl(db)
	productService := service.NewProductServiceImpl(productRepository)
	productHandler := handler.NewProductHandler(productService)

	categoryRepository := repository.NewCategoryRepositoryImpl(db)
	categoryService := service.NewCategoryServiceImpl(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := gin.New()

	r.Use(
		middleware.Logger(),
		middleware.ErrorHandler(),
		middleware.RequestID(),
	)

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUserByEmail)

	api := r.Group("/api")
	api.Use(middleware.JWTAuth())

	api.POST("/products", middleware.JWTAuth(), productHandler.CreateProduct)
	api.GET("/products/:id", middleware.JWTAuth(), productHandler.GetProduct)
	api.GET("/products", middleware.JWTAuth(), productHandler.GetProducts)
	api.PUT("/products", middleware.JWTAuth(), productHandler.UpdateProduct)
	api.DELETE("/products/:id", middleware.JWTAuth(), productHandler.DeleteProduct)

	api.POST("/categories", middleware.JWTAuth(), categoryHandler.CreateCategory)
	api.GET("/categories/:id", middleware.JWTAuth(), categoryHandler.GetCategory)
	api.GET("/categories", middleware.JWTAuth(), categoryHandler.GetCategories)
	api.PUT("/categories", middleware.JWTAuth(), categoryHandler.UpdateCategory)
	api.DELETE("/categories/:id", middleware.JWTAuth(), categoryHandler.DeleteCategory)

	api.PUT("/users", middleware.JWTAuth(), userHandler.UpdateUser)
	api.DELETE("/users", middleware.JWTAuth(), userHandler.DeleteUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed : %v", err)
	}
}
