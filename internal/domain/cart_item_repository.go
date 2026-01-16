package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
)

type CartItemRepository interface {
	Create(ctx context.Context, cartItem *model.CartItem) error
	FindAllByCartId(ctx context.Context, cartId int) ([]model.CartItem, error)
	FindById(ctx context.Context, id int) (model.CartItem, error)
	FindByCartAndProductId(ctx context.Context, cartId int, productId string) (*model.CartItem, error)
	Update(ctx context.Context, id int, quantity int) error
	Delete(ctx context.Context, id int) error
}
