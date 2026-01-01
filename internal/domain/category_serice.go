package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, category *model.Category) error
	GetCategory(ctx context.Context, id string) error
	GetCategories(ctx context.Context) ([]model.Category, error)
	UpdateCategory(ctx context.Context, category *model.Category) error
	DeleteCategory(ctx context.Context, id string) error
}
