package product

import (
	"errors"
	"mini-ecommerce/internal/domain/product"
	"mini-ecommerce/internal/helper"
	"mini-ecommerce/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService product.Service
}

func NewHandler(productService product.Service) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	productData := product.Data{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	if appErr := h.productService.Create(c.Request.Context(), &productData); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Create Product",
		Response{
			ID:          productData.ID,
			CategoryID:  productData.CategoryID,
			Name:        productData.Name,
			Description: productData.Description,
			Price:       productData.Price,
			Stock:       productData.Stock,
		},
	)
	c.JSON(status, res)
}

func (h *ProductHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Product id is required"),
		))
		return
	}

	productData, appErr := h.productService.Get(c.Request.Context(), id)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Get Product",
		Response{
			ID:          productData.ID,
			CategoryID:  productData.CategoryID,
			Name:        productData.Name,
			Description: productData.Description,
			Price:       productData.Price,
			Stock:       productData.Stock,
		},
	)
	c.JSON(status, res)
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, appErr := h.productService.GetAll(c.Request.Context())
	if appErr != nil {
		c.Error(appErr)
		return
	}

	var dataResponses []Response
	for _, product := range products {
		response := Response{
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

func (h *ProductHandler) Update(c *gin.Context) {
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	productUpdate := product.Update{
		ID:          req.ID,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	if appErr := h.productService.Update(c.Request.Context(), &productUpdate); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Update Product",
		Response{
			ID:          productUpdate.ID,
			CategoryID:  *productUpdate.CategoryID,
			Name:        *productUpdate.Name,
			Description: *productUpdate.Description,
			Price:       *productUpdate.Price,
			Stock:       *productUpdate.Stock,
		},
	)
	c.JSON(status, res)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Product id is required"),
		))
		return
	}

	if appErr := h.productService.Delete(c.Request.Context(), id); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Deleted Product")
	c.JSON(status, res)
}
