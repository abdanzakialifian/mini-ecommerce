package order

import "mini-ecommerce/internal/domain/order"

type Response struct {
	ID         int          `json:"id"`
	UserID     int          `json:"user_id"`
	TotalPrice float64      `json:"total_price"`
	Status     order.Status `json:"status"`
}

type ItemResponse struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"oder_id"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type DetailResponse struct {
	Order Response       `json:"order"`
	Items []ItemResponse `json:"items"`
}
