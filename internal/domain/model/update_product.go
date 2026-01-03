package model

import "time"

type UpdateProduct struct {
	ID          string
	CategoryID  *string
	Name        *string
	Description *string
	Price       *float64
	Stock       *int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
