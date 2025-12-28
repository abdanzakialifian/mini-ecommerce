package handler

import (
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/handler/response"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service domain.ProductService
}

func NewProductHandler(service domain.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, appErr := h.service.GetProducts(c.Request.Context())
	if appErr != nil {
		c.Error(appErr)
		return
	}
	status, res := response.Success(
		"Success Get Products",
		toProductResponse(products),
	)
	c.JSON(status, res)
}
