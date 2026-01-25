package product

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, data *Data) error
	Find(ctx context.Context, id string) (Data, error)
	FindAll(ctx context.Context) ([]Data, error)
	Update(ctx context.Context, update *Update) error
	UpdateStock(ctx context.Context, id string, quantity int) error
	Delete(ctx context.Context, id string) error
}
