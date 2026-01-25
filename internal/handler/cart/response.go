package cart

type ItemResponse struct {
	ID        int    `json:"id"`
	CartID    int    `json:"cart_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
