package category

import (
	"errors"
	"mini-ecommerce/internal/domain/category"
	"mini-ecommerce/internal/helper"
	"mini-ecommerce/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service category.CategoryService
}

func NewHandler(service category.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	category := category.Category{Name: req.Name}

	if appErr := h.service.CreateCategory(c.Request.Context(), &category); appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := CategoryResponse{
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

	result, appErr := h.service.GetCategory(c.Request.Context(), id)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := CategoryResponse{
		ID:   result.ID,
		Name: result.Name,
	}

	status, res := response.Success(
		"Success Get Category",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	results, appErr := h.service.GetCategories(c.Request.Context())
	if appErr != nil {
		c.Error(appErr)
		return
	}

	var dataResponses []CategoryResponse
	for _, category := range results {
		response := CategoryResponse{
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
	var req UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	updateCategory := category.UpdateCategory{
		ID:   req.ID,
		Name: req.Name,
	}

	if appErr := h.service.UpdateCategory(c.Request.Context(), &updateCategory); appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := CategoryResponse{
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
