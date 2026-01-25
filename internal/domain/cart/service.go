package cart

import (
	"context"
	"mini-ecommerce/internal/helper"
)

type Service interface {
	GetItems(ctx context.Context, userId int) ([]Item, *helper.AppError)
	AddItem(ctx context.Context, userId int, productId string, quantity int) (Item, *helper.AppError)
	UpdateItemQuantity(ctx context.Context, userId int, updateItem UpdateItem) *helper.AppError
	DeleteItem(ctx context.Context, userId int, itemId int) *helper.AppError
}
