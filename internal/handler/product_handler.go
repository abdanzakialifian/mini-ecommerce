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

type ProductHandler struct {
	service domain.ProductService
}

func NewProductHandler(service domain.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req request.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	product := model.Product{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if appErr := h.service.CreateProduct(c.Request.Context(), &product); appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := response.ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}

	status, res := response.Success(
		"Success Create Product",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Product id is required"),
		))
		return
	}

	product, appErr := h.service.GetProduct(c.Request.Context(), id)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := response.ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}

	status, res := response.Success(
		"Success Get Product",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, appErr := h.service.GetProducts(c.Request.Context())
	if appErr != nil {
		c.Error(appErr)
		return
	}

	var dataResponses []response.ProductResponse
	for _, product := range products {
		response := response.ProductResponse{
			ID:          product.ID,
			CategoryID:  product.CategoryID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		}
		dataResponses = append(dataResponses, response)
	}

	status, res := response.Success(
		"Success Get Products",
		dataResponses,
	)
	c.JSON(status, res)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var req request.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	updateProduct := model.UpdateProduct{
		ID:          req.ID,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if appErr := h.service.UpdateProduct(c.Request.Context(), &updateProduct); appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := response.ProductResponse{
		ID:          updateProduct.ID,
		CategoryID:  *updateProduct.CategoryID,
		Name:        *updateProduct.Name,
		Description: *updateProduct.Description,
		Price:       *updateProduct.Price,
		Stock:       *updateProduct.Stock,
	}

	status, res := response.Success(
		"Success Update Product",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Product id is required"),
		))
		return
	}

	if appErr := h.service.DeleteProduct(c.Request.Context(), id); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Deleted Product")
	c.JSON(status, res)
}
