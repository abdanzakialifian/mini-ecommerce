package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *model.Category) error
	Find(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]model.Category, error)
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id string) error
}
