package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	FindAll(ctx context.Context) ([]model.Product, error)
}
