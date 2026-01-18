package category

import (
	"context"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	Find(ctx context.Context, id string) (Category, error)
	FindAll(ctx context.Context) ([]Category, error)
	Update(ctx context.Context, updateCategory *UpdateCategory) error
	Delete(ctx context.Context, id string) error
}
