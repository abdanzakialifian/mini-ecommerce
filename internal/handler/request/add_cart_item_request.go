package request

type AddCartItemRequest struct {
	ProductId string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}
