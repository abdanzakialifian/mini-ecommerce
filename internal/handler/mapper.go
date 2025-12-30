package handler

import (
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/handler/request"
	"mini-ecommerce/internal/handler/response"
)

func toProductResponses(products []model.Product) []response.ProductResponse {
	var responses []response.ProductResponse
	for _, product := range products {
		response := response.ProductResponse{
			ID:          product.ID,
			CategoryID:  product.CategoryID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
		responses = append(responses, response)
	}
	return responses
}

func toProductResponse(product model.Product) response.ProductResponse {
	return response.ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func toProduct(product request.CreateProductRequest) model.Product {
	return model.Product{
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
