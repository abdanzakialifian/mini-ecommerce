package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]model.Product, error)
}
