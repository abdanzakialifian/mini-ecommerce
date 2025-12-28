package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
)

type ProductService interface {
	GetProducts(ctx context.Context) ([]model.Product, error)
}
