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

	productRepository := repository.NewProduct(db, tx)
	productService := service.NewProduct(productRepository)
	productHandler := product.NewHandler(productService)

	categoryRepository := repository.NewCategory(db)
	categoryService := service.NewCategory(categoryRepository)
	categoryHandler := category.NewHandler(categoryService)

	userRepository := repository.NewUser(db)
	userService := service.NewUser(userRepository)
	userHandler := user.NewHandler(userService)

	cartRepository := repository.NewCart(tx)
	cartItemRepository := repository.NewCartItem(tx)
	cartService := service.NewCart(tx, cartRepository, cartItemRepository)
	cartHandler := cart.NewHandler(cartService)

	r := gin.New()

	r.Use(
		middleware.Logger(),
		middleware.ErrorHandler(),
		middleware.RequestID(),
	)

	// ======================== without token ========================
	r.POST("/users", userHandler.Create)
	r.GET("/users", userHandler.GetByEmail)

	// ======================== with token ========================
	api := r.Group("/api")
	api.Use(middleware.JWTAuth())

	api.POST("/products", productHandler.Create)
	api.GET("/products/:id", productHandler.Get)
	api.GET("/products", productHandler.GetAll)
	api.PUT("/products", productHandler.Update)
	api.DELETE("/products/:id", productHandler.Delete)

	api.POST("/categories", categoryHandler.Create)
	api.GET("/categories/:id", categoryHandler.Get)
	api.GET("/categories", categoryHandler.GetAll)
	api.PUT("/categories", categoryHandler.Update)
	api.DELETE("/categories/:id", categoryHandler.Delete)

	api.PUT("/users", userHandler.Update)
	api.DELETE("/users", userHandler.Delete)

	api.POST("/carts", cartHandler.AddItem)
	api.GET("/carts", cartHandler.GetItems)
	api.PUT("/carts", cartHandler.UpdateItemQuantity)
	api.DELETE("/carts/:cart_item_id", cartHandler.DeleteItem)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed : %v", err)
	}
}
