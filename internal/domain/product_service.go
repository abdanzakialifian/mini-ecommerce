package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"
)

type ProductService interface {
	GetProducts(ctx context.Context) ([]model.Product, *helper.AppError)
}
