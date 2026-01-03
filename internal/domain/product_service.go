package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *model.Product) *helper.AppError
	GetProduct(ctx context.Context, id string) (model.Product, *helper.AppError)
	GetProducts(ctx context.Context) ([]model.Product, *helper.AppError)
	UpdateProduct(ctx context.Context, product model.Product) (model.Product, *helper.AppError)
	DeleteProduct(ctx context.Context, id string) *helper.AppError
}
