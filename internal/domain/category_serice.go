package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, category *model.Category) *helper.AppError
	GetCategory(ctx context.Context, id string) (model.Category, *helper.AppError)
	GetCategories(ctx context.Context) ([]model.Category, *helper.AppError)
	UpdateCategory(ctx context.Context, category *model.Category) *helper.AppError
	DeleteCategory(ctx context.Context, id string) *helper.AppError
}
