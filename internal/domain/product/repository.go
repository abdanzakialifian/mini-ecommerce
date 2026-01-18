package product

import (
	"context"
)

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	Find(ctx context.Context, id string) (Product, error)
	FindAll(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, updateProduct *UpdateProduct) error
	Delete(ctx context.Context, id string) error
}
