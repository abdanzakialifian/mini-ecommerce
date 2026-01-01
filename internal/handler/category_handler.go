package handler

import (
	"mini-ecommerce/internal/domain"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service domain.CategoryService
}

func NewCategoryHandler(service domain.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	panic("unimplemented")
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
	panic("unimplemented")
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	panic("unimplemented")
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	panic("unimplemented")
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	panic("unimplemented")
}
