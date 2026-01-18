package cart

import (
	"errors"
	"mini-ecommerce/internal/domain/cart"
	"mini-ecommerce/internal/helper"
	"mini-ecommerce/internal/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service cart.CartService
}

func NewHandler(service cart.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) AddCartItemToCart(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	var req AddCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	result, appErr := h.service.AddCartItemToCart(c.Request.Context(), userId, req.ProductId, req.Quantity)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := CartItemResponse{
		ID:        result.ID,
		CartID:    result.CartID,
		ProductID: result.ProductID,
		Quantity:  result.Quantity,
	}

	status, res := response.Success(
		"Success Add Cart Item",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *CartHandler) GetCartItems(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	results, appErr := h.service.GetCartItems(c.Request.Context(), userId)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	var dataResponses []CartItemResponse
	for _, cartItem := range results {
		response := CartItemResponse{
			ID:        cartItem.ID,
			CartID:    cartItem.CartID,
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
		}
		dataResponses = append(dataResponses, response)
	}

	status, res := response.Success(
		"Success Get Cart Items",
		dataResponses,
	)
	c.JSON(status, res)
}

func (h *CartHandler) UpdateCartItemQuantity(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	var req UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	updateCartItem := cart.UpdateCartItem{
		ID:       req.CartItemId,
		Quantity: req.Quantity,
	}
	if appErr := h.service.UpdateCartItemQuantity(c.Request.Context(), userId, updateCartItem); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Update Quantity")
	c.JSON(status, res)
}

func (h *CartHandler) DeleteCartItemFromCart(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	id := c.Param("cart_item_id")
	if id == "" {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			errors.New("Cart item id is required"),
		))
		return
	}

	cartItemId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Internal Server Error",
			errors.New("Error convert cart item id"),
		))
		return
	}

	if appErr := h.service.DeleteCartItemFromCart(c.Request.Context(), userId, cartItemId); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Delete Cart Item")
	c.JSON(status, res)
}
