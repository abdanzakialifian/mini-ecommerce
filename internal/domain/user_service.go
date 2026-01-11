package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
	"mini-ecommerce/internal/helper"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) *helper.AppError
	GetUser(ctx context.Context, login model.LoginUser) (model.User, string, *helper.AppError)
	UpdateUser(ctx context.Context, updateUser *model.UpdateUser) *helper.AppError
	DeleteUser(ctx context.Context, id int) *helper.AppError
}
