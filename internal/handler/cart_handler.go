package handler

import (
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/handler/request"
	"mini-ecommerce/internal/handler/response"
	"mini-ecommerce/internal/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service domain.CartService
}

func NewCartHandler(service domain.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) AddCartItemToCart(c *gin.Context) {
	userId := c.MustGet("user_id").(int)
	var req request.AddCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	cartItem, err := h.service.AddCartItemToCart(c.Request.Context(), userId, req.ProductId, req.Quantity)
	if err != nil {
		c.Error(err)
		return
	}

	status, res := response.Success(
		"Success Add Cart Item",
		toCartItemResponse(cartItem),
	)
	c.JSON(status, res)
}

func (h *CartHandler) GetCartItems(c *gin.Context) {
	userId := c.MustGet("user_id").(int)
	cartItems, err := h.service.GetCartItems(c.Request.Context(), userId)
	if err != nil {
		c.Error(err)
		return
	}

	status, res := response.Success(
		"Success Get Cart Items",
		toCartItemsResponse(cartItems),
	)
	c.JSON(status, res)
}

func (h *CartHandler) UpdateCartItemQuantity(c *gin.Context) {
	userId := c.MustGet("user_id").(int)
	var req request.UpdateQuantityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}
	if err := h.service.UpdateCartItemQuantity(c.Request.Context(), userId, req.CartItemId, req.Quantity); err != nil {
		c.Error(err)
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

	if err := h.service.DeleteCartItemFromCart(c.Request.Context(), userId, cartItemId); err != nil {
		c.Error(err)
		return
	}

	status, res := response.SuccessNoContent("Success Delete Cart Item")
	c.JSON(status, res)
}
