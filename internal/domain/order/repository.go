package order

import "context"

type Repository interface {
	Create(ctx context.Context, data *Data) error
	FindById(ctx context.Context, id int) (Data, error)
	FindByUserId(ctx context.Context, userId int) ([]Data, error)
	Update(ctx context.Context, update *Update) error
	UpdateStatus(ctx context.Context, id int, status Status) error
	Delete(ctx context.Context, id int) error
}

type ItemRepository interface {
	CreateItems(ctx context.Context, items []Item) error
	FindItems(ctx context.Context, orderId int) ([]Item, error)
}
