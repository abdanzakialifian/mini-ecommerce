package main

import (
	"log"
	"mini-ecommerce/internal/database"
	"mini-ecommerce/internal/handler"
	"mini-ecommerce/internal/repository"
	"mini-ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	productRepository := repository.NewProductRepositoryImpl(db)
	productService := service.NewProductServiceImpl(productRepository)
	productHandler := handler.NewProductHandler(productService)

	r := gin.Default()
	r.GET("/products", productHandler.GetProducts)
	r.Run(":8080")
}
