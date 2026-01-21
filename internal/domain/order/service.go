package order

import (
	"context"
	"mini-ecommerce/internal/helper"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userId int, createOrderItems []CreateOrderItem) (OrderDetail, *helper.AppError)
	GetOrder(ctx context.Context, id int) (OrderDetail, *helper.AppError)
	GetOrderByUserId(ctx context.Context, userId int) ([]OrderDetail, *helper.AppError)
	UpdateOrderStatus(ctx context.Context, id int, status UpdateOrder) *helper.AppError
	CancelOrder(ctx context.Context, id int) *helper.AppError
}
