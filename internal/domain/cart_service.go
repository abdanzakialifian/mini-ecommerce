package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"
)

type CartService interface {
	GetCartItems(ctx context.Context, userId int) ([]model.CartItem, *helper.AppError)
	AddCartItemToCart(ctx context.Context, userId int, productId string, quantity int) (model.CartItem, *helper.AppError)
	UpdateCartItemQuantity(ctx context.Context, userId int, updateCartItem model.UpdateCartItem) *helper.AppError
	DeleteCartItemFromCart(ctx context.Context, userId int, cartItemId int) *helper.AppError
}
