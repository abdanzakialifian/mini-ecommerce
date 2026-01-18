package category

import (
	"context"
	"mini-ecommerce/internal/helper"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, category *Category) *helper.AppError
	GetCategory(ctx context.Context, id string) (Category, *helper.AppError)
	GetCategories(ctx context.Context) ([]Category, *helper.AppError)
	UpdateCategory(ctx context.Context, updateCategory *UpdateCategory) *helper.AppError
	DeleteCategory(ctx context.Context, id string) *helper.AppError
}
