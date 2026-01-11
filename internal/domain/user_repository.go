package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, login model.LoginUser) (model.User, string, error)
	FindById(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, updateUser *model.UpdateUser) error
	Delete(ctx context.Context, id int) error
}
