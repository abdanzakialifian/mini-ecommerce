package model

type UpdateProduct struct {
	ID          string
	CategoryID  *string
	Name        *string
	Description *string
	Price       *float64
	Stock       *int
}
