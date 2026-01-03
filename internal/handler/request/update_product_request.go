package request

type UpdateProductRequest struct {
	ID          string   `json:"id" binding:"required"`
	CategoryID  *string  `json:"category_id,omitempty"`
	Name        *string  `json:"name" binding:"omitempty,min=3,max=50"`
	Description *string  `json:"description" binding:"omitempty,max=255"`
	Price       *float64 `json:"price,omitempty"`
	Stock       *int     `json:"stock,omitempty"`
}
