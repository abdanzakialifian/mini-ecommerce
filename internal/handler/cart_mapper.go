package handler

import (
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/handler/response"
)

func toCartItemResponse(cartItem model.CartItem) response.CartItemResponse {
	return response.CartItemResponse{
		ID:        cartItem.ID,
		CartID:    cartItem.CartID,
		ProductID: cartItem.ProductID,
		Quantity:  cartItem.Quantity,
	}
}

func toCartItemsResponse(cartItems []model.CartItem) []response.CartItemResponse {
	var cartItemsResponse []response.CartItemResponse
	for _, cartItem := range cartItems {
		response := response.CartItemResponse{
			ID:        cartItem.ID,
			CartID:    cartItem.CartID,
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
		}
		cartItemsResponse = append(cartItemsResponse, response)
	}
	return cartItemsResponse
}
