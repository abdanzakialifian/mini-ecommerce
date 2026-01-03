package handler

import (
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/handler/request"
	"mini-ecommerce/internal/handler/response"
)

func ToCategoryFromCreate(req request.CreateCategoryRequest) model.Category {
	return model.Category{
		Name: req.Name,
	}
}

func ToCategoryFromUpdate(req request.UpdateCategoryRequest) model.Category {
	return model.Category{
		ID:   req.ID,
		Name: req.Name,
	}
}

func ToCategoryResponse(category model.Category) response.CategoryResponse {
	return response.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []model.Category) []response.CategoryResponse {
	var responses []response.CategoryResponse
	for _, category := range categories {
		response := response.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		}
		responses = append(responses, response)
	}
	return responses
}
