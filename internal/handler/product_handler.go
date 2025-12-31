package handler

import (
	"errors"
	"mini-ecommerce/internal/domain"
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

	product := toProductFromCreate(req)
	if err := h.service.CreateProduct(c.Request.Context(), &product); err != nil {
		c.Error(err)
		return
	}

	status, res := response.Success(
		"Success Create Product",
		toProductResponse(product),
	)
	c.JSON(status, res)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, appErr := h.service.GetProducts(c.Request.Context())
	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Get Products",
		toProductResponses(products),
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

	product := toProductFromUpdate(req)
	if err := h.service.UpdateProduct(c.Request.Context(), &product); err != nil {
		c.Error(err)
		return
	}

	status, res := response.Success(
		"Success Update Product",
		toProductResponse(product),
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

	err := h.service.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	status, res := response.SuccessNoContent("Success Deleted Product")
	c.JSON(status, res)
}
