package handler

import (
	"errors"
	"mini-ecommerce/internal/domain"
	"mini-ecommerce/internal/domain/model"
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

	cartItem, appErr := h.service.AddCartItemToCart(c.Request.Context(), userId, req.ProductId, req.Quantity)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	dataResponse := response.CartItemResponse{
		ID:        cartItem.ID,
		CartID:    cartItem.CartID,
		ProductID: cartItem.ProductID,
		Quantity:  cartItem.Quantity,
	}

	status, res := response.Success(
		"Success Add Cart Item",
		dataResponse,
	)
	c.JSON(status, res)
}

func (h *CartHandler) GetCartItems(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	cartItems, appErr := h.service.GetCartItems(c.Request.Context(), userId)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	var dataResponses []response.CartItemResponse
	for _, cartItem := range cartItems {
		response := response.CartItemResponse{
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

	var req request.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	updateCartItem := model.UpdateCartItem{
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
