package handler

import (
	"mini-ecommerce/internal/domain"
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get products",
		})
		return
	}
	c.JSON(http.StatusOK, toProductResponse(products))
}
