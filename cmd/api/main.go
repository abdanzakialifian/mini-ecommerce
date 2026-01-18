package main

import (
	"context"
	"log"
	"mini-ecommerce/internal/database"
	"mini-ecommerce/internal/handler/cart"
	"mini-ecommerce/internal/handler/category"
	"mini-ecommerce/internal/handler/product"
	"mini-ecommerce/internal/handler/user"
	"mini-ecommerce/internal/helper"
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

	tx := helper.NewTransaction(db)

	productRepository := repository.NewProductRepositoryImpl(db)
	productService := service.NewProductServiceImpl(productRepository)
	productHandler := product.NewHandler(productService)

	categoryRepository := repository.NewCategoryRepositoryImpl(db)
	categoryService := service.NewCategoryServiceImpl(categoryRepository)
	categoryHandler := category.NewHandler(categoryService)

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userHandler := user.NewHandler(userService)

	cartRepository := repository.NewCartRepositoryImpl(tx)
	cartItemRepository := repository.NewCartItemRepositoryImpl(tx)
	cartService := service.NewCartServiceImpl(tx, cartRepository, cartItemRepository)
	cartHandler := cart.NewHandler(cartService)

	r := gin.New()

	r.Use(
		middleware.Logger(),
		middleware.ErrorHandler(),
		middleware.RequestID(),
	)

	// ======================== without token ========================
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUserByEmail)

	// ======================== with token ========================
	api := r.Group("/api")
	api.Use(middleware.JWTAuth())

	api.POST("/products", productHandler.CreateProduct)
	api.GET("/products/:id", productHandler.GetProduct)
	api.GET("/products", productHandler.GetProducts)
	api.PUT("/products", productHandler.UpdateProduct)
	api.DELETE("/products/:id", productHandler.DeleteProduct)

	api.POST("/categories", categoryHandler.CreateCategory)
	api.GET("/categories/:id", categoryHandler.GetCategory)
	api.GET("/categories", categoryHandler.GetCategories)
	api.PUT("/categories", categoryHandler.UpdateCategory)
	api.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	api.PUT("/users", userHandler.UpdateUser)
	api.DELETE("/users", userHandler.DeleteUser)

	api.POST("/carts", cartHandler.AddCartItemToCart)
	api.GET("/carts", cartHandler.GetCartItems)
	api.PUT("/carts", cartHandler.UpdateCartItemQuantity)
	api.DELETE("/carts/:cart_item_id", cartHandler.DeleteCartItemFromCart)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed : %v", err)
	}
}
