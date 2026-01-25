package cart

import "context"

type Repository interface {
	FindByUserId(ctx context.Context, userId int) (Data, error)
	FindOrCreateByUserId(ctx context.Context, userId int) (Data, error)
}

type ItemRepository interface {
	Create(ctx context.Context, item *Item) error
	FindAllByCartId(ctx context.Context, cartId int) ([]Item, error)
	FindById(ctx context.Context, itemId int) (Item, error)
	FindByCartAndProductId(ctx context.Context, cartId int, productId string) (*Item, error)
	Update(ctx context.Context, updateItem UpdateItem) error
	Delete(ctx context.Context, itemId int) error
}
