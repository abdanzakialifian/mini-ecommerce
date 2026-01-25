package order

import (
	"context"
	"mini-ecommerce/internal/helper"
)

type Service interface {
	Create(ctx context.Context, userId int, newItems []NewItem) (Detail, *helper.AppError)
	Get(ctx context.Context, id int) (Detail, *helper.AppError)
	GetByUserId(ctx context.Context, userId int) ([]Detail, *helper.AppError)
	UpdateStatus(ctx context.Context, id int, status Status) *helper.AppError
	Cancel(ctx context.Context, id int) *helper.AppError
}
