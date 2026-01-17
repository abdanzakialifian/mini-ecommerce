package handler

import (
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/handler/request"
	"mini-ecommerce/internal/handler/response"
	"mini-ecommerce/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service domain.CategoryService
}

func NewCategoryHandler(service domain.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	category := model.Category{Name: req.Name}

	if appErr := h.service.CreateCategory(c.Request.Context(), &category); appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := response.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	status, res := response.Success(
		"Success Create Category",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Category id is required"),
		))
		return
	}

	category, appErr := h.service.GetCategory(c.Request.Context(), id)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := response.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	status, res := response.Success(
		"Success Get Category",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	categories, appErr := h.service.GetCategories(c.Request.Context())
	if appErr != nil {
		c.Error(appErr)
		return
	}

	var dataResponses []response.CategoryResponse
	for _, category := range categories {
		response := response.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		}
		dataResponses = append(dataResponses, response)
	}

	status, res := response.Success(
		"Success Get Categories",
		dataResponses,
	)
	c.JSON(status, res)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	var req request.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	updateCategory := model.UpdateCategory{
		ID:   req.ID,
		Name: req.Name,
	}

	if appErr := h.service.UpdateCategory(c.Request.Context(), &updateCategory); appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := response.CategoryResponse{
		ID:   updateCategory.ID,
		Name: updateCategory.Name,
	}

	status, res := response.Success(
		"Success Update Category",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Category id is required"),
		))
		return
	}

	if appErr := h.service.DeleteCategory(c.Request.Context(), id); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Delete Category")
	c.JSON(status, res)
}
