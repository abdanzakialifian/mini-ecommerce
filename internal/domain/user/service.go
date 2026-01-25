package user

import (
	"context"
	"mini-ecommerce/internal/helper"
)

type Service interface {
	Create(ctx context.Context, data *Data) *helper.AppError
	GetByEmail(ctx context.Context, login Login) (Data, string, *helper.AppError)
	Update(ctx context.Context, update *Update) *helper.AppError
	Delete(ctx context.Context, id int) *helper.AppError
}
