package cart

import "context"

type CartRepository interface {
	FindByUserId(ctx context.Context, userId int) (Cart, error)
	FindOrCreateByUserId(ctx context.Context, userId int) (Cart, error)
}

type CartItemRepository interface {
	Create(ctx context.Context, cartItem *CartItem) error
	FindAllByCartId(ctx context.Context, cartId int) ([]CartItem, error)
	FindById(ctx context.Context, id int) (CartItem, error)
	FindByCartAndProductId(ctx context.Context, cartId int, productId string) (*CartItem, error)
	Update(ctx context.Context, updateCartItem UpdateCartItem) error
	Delete(ctx context.Context, id int) error
}
