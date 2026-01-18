package product

import (
	"context"
	"mini-ecommerce/internal/helper"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *Product) *helper.AppError
	GetProduct(ctx context.Context, id string) (Product, *helper.AppError)
	GetProducts(ctx context.Context) ([]Product, *helper.AppError)
	UpdateProduct(ctx context.Context, updateProduct *UpdateProduct) *helper.AppError
	DeleteProduct(ctx context.Context, id string) *helper.AppError
}
