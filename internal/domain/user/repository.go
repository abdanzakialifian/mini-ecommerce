package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, data *Data) error
	FindByEmail(ctx context.Context, login Login) (Data, string, error)
	FindById(ctx context.Context, id int) (Data, error)
	Update(ctx context.Context, update *Update) error
	Delete(ctx context.Context, id int) error
}
