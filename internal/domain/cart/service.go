package cart

import (
	"context"
	"mini-ecommerce/internal/helper"
)

type CartService interface {
	GetCartItems(ctx context.Context, userId int) ([]CartItem, *helper.AppError)
	AddCartItemToCart(ctx context.Context, userId int, productId string, quantity int) (CartItem, *helper.AppError)
	UpdateCartItemQuantity(ctx context.Context, userId int, updateCartItem UpdateCartItem) *helper.AppError
	DeleteCartItemFromCart(ctx context.Context, userId int, cartItemId int) *helper.AppError
}
