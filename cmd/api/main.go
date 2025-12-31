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
)

func main() {
	ctx := context.Background()

	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect db m: %v", err)
	}
	defer db.Close()

	productRepository := repository.NewProductRepositoryImpl(db)
	productService := service.NewProductServiceImpl(productRepository)
	productHandler := handler.NewProductHandler(productService)

	r := gin.New()

	r.Use(
		middleware.Logger(),
		middleware.ErrorHandler(),
		middleware.RequestID(),
	)

	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products", productHandler.GetProducts)
	r.PUT("/products", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed : %v", err)
	}
}
