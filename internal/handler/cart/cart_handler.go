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
	cartService cart.Service
}

func NewHandler(cartService cart.Service) *CartHandler {
	return &CartHandler{cartService: cartService}
}

func (h *CartHandler) AddItem(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	var req AddItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	cartItem, appErr := h.cartService.AddItem(c.Request.Context(), userId, req.ProductId, req.Quantity)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.Success(
		"Success Add Cart Item",
		ItemResponse{
			ID:        cartItem.ID,
			CartID:    cartItem.CartID,
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
		},
	)
	c.JSON(status, res)
}

func (h *CartHandler) GetItems(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	cartItems, appErr := h.cartService.GetItems(c.Request.Context(), userId)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	var itemResponses []ItemResponse
	for _, cartItem := range cartItems {
		itemResponse := ItemResponse{
			ID:        cartItem.ID,
			CartID:    cartItem.CartID,
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
		}
		itemResponses = append(itemResponses, itemResponse)
	}

	status, res := response.Success(
		"Success Get Cart Items",
		itemResponses,
	)
	c.JSON(status, res)
}

func (h *CartHandler) UpdateItemQuantity(c *gin.Context) {
	userId := c.MustGet("user_id").(int)

	var req UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(helper.NewAppError(
			http.StatusBadRequest,
			"Invalid Request Body",
			err,
		))
		return
	}

	updateItem := cart.UpdateItem{
		ID:       req.CartItemId,
		Quantity: req.Quantity,
	}
	if appErr := h.cartService.UpdateItemQuantity(c.Request.Context(), userId, updateItem); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Update Quantity")
	c.JSON(status, res)
}

func (h *CartHandler) DeleteItem(c *gin.Context) {
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
			http.StatusInternalServerError,
			"Internal Server Error",
			errors.New("Error convert cart item id"),
		))
		return
	}

	if appErr := h.cartService.DeleteItem(c.Request.Context(), userId, cartItemId); appErr != nil {
		c.Error(appErr)
		return
	}

	status, res := response.SuccessNoContent("Success Delete Cart Item")
	c.JSON(status, res)
}
