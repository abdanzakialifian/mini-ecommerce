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
	categoryService category.Service
}

func NewHandler(categoryService category.Service) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	categoryData := category.Data{Name: req.Name}
	if appErr := h.categoryService.Create(c.Request.Context(), &categoryData); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Create Category",
		Response{
			ID:   categoryData.ID,
			Name: categoryData.Name,
		},
	)
	c.JSON(status, res)
}

func (h *CategoryHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Category id is required"),
		))
		return
	}

	categoryData, appErr := h.categoryService.Get(c.Request.Context(), id)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Get Category",
		Response{
			ID:   categoryData.ID,
			Name: categoryData.Name,
		},
	)
	c.JSON(status, res)
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, appErr := h.categoryService.GetAll(c.Request.Context())
	if appErr != nil {
		c.Error(appErr)
		return
	}

	var categoryResponses []Response
	for _, category := range categories {
		response := Response{
			ID:   category.ID,
			Name: category.Name,
		}
		categoryResponses = append(categoryResponses, response)
	}

	status, res := response.Success(
		"Success Get Categories",
		categoryResponses,
	)
	c.JSON(status, res)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	categoryUpdate := category.Update{
		ID:   req.ID,
		Name: req.Name,
	}
	if appErr := h.categoryService.Update(c.Request.Context(), &categoryUpdate); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Update Category",
		Response{
			ID:   categoryUpdate.ID,
			Name: categoryUpdate.Name,
		},
	)
	c.JSON(status, res)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Category id is required"),
		))
		return
	}

	if appErr := h.categoryService.Delete(c.Request.Context(), id); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Delete Category")
	c.JSON(status, res)
}
