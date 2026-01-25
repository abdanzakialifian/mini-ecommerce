package order

import "mini-ecommerce/internal/domain/order"

type CreateRequest struct {
	Items []ItemRequest `json:"items" binding:"required"`
}

type ItemRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type UpdateStatusRequest struct {
	Status order.Status `json:"status" binding:"required"`
}
