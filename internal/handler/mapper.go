package handler

import (
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/handler/response"
)

func toProductResponse(products []model.Product) []response.ProductResponse {
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
