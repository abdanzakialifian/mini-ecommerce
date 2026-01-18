package product

type CreateProductRequest struct {
	CategoryID  string  `json:"category_id" binding:"required,gt=0"`
	Name        string  `json:"name" binding:"required,min=3,max=50"`
	Description string  `json:"description" binding:"omitempty,max=255"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
}

type UpdateProductRequest struct {
	ID          string   `json:"id" binding:"required"`
	CategoryID  *string  `json:"category_id,omitempty"`
	Name        *string  `json:"name" binding:"omitempty,min=3,max=50"`
	Description *string  `json:"description" binding:"omitempty,max=255"`
	Price       *float64 `json:"price,omitempty"`
	Stock       *int     `json:"stock,omitempty"`
}
