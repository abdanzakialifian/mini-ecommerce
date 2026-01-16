package domain

import (
	"context"
	"mini-ecommerce/internal/domain/model"
)

type CartRepository interface {
	FindByUserId(ctx context.Context, userId int) (model.Cart, error)
	FindOrCreateByUserId(ctx context.Context, userId int) (model.Cart, error)
}
