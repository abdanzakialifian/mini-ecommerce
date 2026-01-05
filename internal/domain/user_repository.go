package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Find(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, user model.User) (model.User, error)
	Delete(ctx context.Context, id int) error
}
