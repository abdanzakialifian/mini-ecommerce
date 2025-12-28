package handler

import (
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/handler/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service domain.ProductService
}

func NewProductHandler(service domain.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetProducts(c.Request.Context())
	if err != nil {
		status, res := response.Error(
			"Failed Get Products",
			err.Error(),
			http.StatusInternalServerError,
		)
		c.JSON(status, res)
		return
	}
	status, res := response.Success(
		"Success Get Products",
		toProductResponse(products),
	)
	c.JSON(status, res)
}
