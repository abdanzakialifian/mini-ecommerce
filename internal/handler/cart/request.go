package cart

type AddItemRequest struct {
	ProductId string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type UpdateItemRequest struct {
	CartItemId int `json:"cart_item_id" binding:"required"`
	Quantity   int `json:"quantity" binding:"required,min=1"`
}
