package category

import (
	"context"
	"mini-ecommerce/internal/helper"
)

type Service interface {
	Create(ctx context.Context, data *Data) *helper.AppError
	Get(ctx context.Context, id string) (Data, *helper.AppError)
	GetAll(ctx context.Context) ([]Data, *helper.AppError)
	Update(ctx context.Context, update *Update) *helper.AppError
	Delete(ctx context.Context, id string) *helper.AppError
}
