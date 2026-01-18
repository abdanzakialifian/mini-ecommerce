package order

import "context"

type OrderRepository interface {
	Create(ctx context.Context, order *Order) error
	FindById(ctx context.Context, id int) (Order, error)
	FindByUserId(ctx context.Context, userId int) ([]Order, error)
	Update(ctx context.Context, updateOrder *UpdateOrder) error
	UpdateStatus(ctx context.Context, id int, status Status) (Order, error)
	Delete(ctx context.Context, id int) error
}

type OrderItemRepository interface {
	CreateOrderItems(ctx context.Context, items []OrderItem) error
	FindOrderItems(ctx context.Context, orderId int) ([]OrderItem, error)
}
