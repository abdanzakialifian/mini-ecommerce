package response

import "time"

type UpdateProductResponse struct {
	ID          string    `json:"id"`
	CategoryID  *string   `json:"category_id,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Price       *float64  `json:"price,omitempty"`
	Stock       *int      `json:"stock,omitempty"`
	UpdatedAt   time.Time `json:"updated_at"`
}
