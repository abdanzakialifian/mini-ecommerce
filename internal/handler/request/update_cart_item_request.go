package request

type UpdateCartItemRequest struct {
	CartItemId int `json:"cart_item_id" binding:"required"`
	Quantity   int `json:"quantity" binding:"required,min=1"`
}
