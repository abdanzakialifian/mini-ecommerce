package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	Find(ctx context.Context, id string) (model.Product, error)
	FindAll(ctx context.Context) ([]model.Product, error)
	Update(ctx context.Context, product model.Product) (model.Product, error)
	Delete(ctx context.Context, id string) error
}
