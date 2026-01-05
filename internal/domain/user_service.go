package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) *helper.AppError
	GetUser(ctx context.Context, id int) (model.User, *helper.AppError)
	UpdateUser(ctx context.Context, user model.User) (model.User, *helper.AppError)
	DeleteUser(ctx context.Context, id int) *helper.AppError
}
