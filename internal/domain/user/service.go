package user

import (
	"context"
	"mini-ecommerce/internal/helper"
)

type UserService interface {
	CreateUser(ctx context.Context, user *User) *helper.AppError
	GetUserByEmail(ctx context.Context, login LoginUser) (User, string, *helper.AppError)
	UpdateUser(ctx context.Context, updateUser *UpdateUser) *helper.AppError
	DeleteUser(ctx context.Context, id int) *helper.AppError
}
