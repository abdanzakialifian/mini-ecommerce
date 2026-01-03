package response

import "time"

type ProductResponse struct {
	ID          string    `json:"id"`
	CategoryID  string    `json:"category_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Stock       int       `json:"stock,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
