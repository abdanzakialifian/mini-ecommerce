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
			CategoryID:  *product.CategoryID,
			Name:        *product.Name,
			Description: *product.Description,
			Price:       *product.Price,
			Stock:       *product.Stock,
		}
		responses = append(responses, response)
	}
	return responses
}

func toProductResponse(product model.Product) response.ProductResponse {
	return response.ProductResponse{
		ID:          product.ID,
		CategoryID:  *product.CategoryID,
		Name:        *product.Name,
		Description: *product.Description,
		Price:       *product.Price,
		Stock:       *product.Stock,
	}
}

func toProductFromCreate(req request.CreateProductRequest) model.Product {
	return model.Product{
		CategoryID:  &req.CategoryID,
		Name:        &req.Name,
		Description: &req.Description,
		Price:       &req.Price,
		Stock:       &req.Stock,
	}
}

func toProductFromUpdate(req request.UpdateProductRequest) model.Product {
	return model.Product{
		ID:          req.ID,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}
}
