package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, login LoginUser) (User, string, error)
	FindById(ctx context.Context, id int) (User, error)
	Update(ctx context.Context, updateUser *UpdateUser) error
	Delete(ctx context.Context, id int) error
}
